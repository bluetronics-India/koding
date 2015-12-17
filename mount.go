package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/koding/klientctl/klientctlerrors"
	"github.com/koding/klientctl/util"
)

// MountCommand mounts a folder on remote machine to local folder by machine
// name.
func MountCommand(c *cli.Context) int {
	if len(c.Args()) < 2 {
		cli.ShowCommandHelp(c, "mount")
		return 1
	}

	var (
		name       = c.Args()[0]
		localPath  = c.Args()[1]
		remotePath = c.String("remotepath") // note the lowercase of all chars
		noIgnore   = c.Bool("noignore")     // note the lowercase of all chars
		noPrefetch = c.Bool("noprefetch")   // note the lowercase of all chars
		noWatch    = c.Bool("nowatch")      // note the lowercase of all chars
	)

	// allow scp like declaration, ie `<machine name>:/path/to/remote`
	if strings.Contains(name, ":") {
		names := strings.Split(name, ":")
		name, remotePath = names[0], names[1]
	}

	// send absolute local path to klient unless local path is empty
	if strings.TrimSpace(localPath) != "" {
		absoluteLocalPath, err := filepath.Abs(localPath)
		if err == nil {
			localPath = absoluteLocalPath
		}
	}

	// remove trailing slashes in remote argument
	remotePath = path.Clean(remotePath)

	// Ask the user if they want the localPath created, if it does not exist.
	if err := askToCreate(localPath, os.Stdin, os.Stdout); err != nil {
		// If the error is that the user cancelled, just return
		if err == klientctlerrors.ErrUserCancelled {
			fmt.Println("Cannot mount to a directory that does not exist, exiting..")
			return 0
		}

		fmt.Printf("Error: Unable to create specified localPath '%s'", localPath)
		return 1
	}

	k, err := CreateKlientClient(NewKlientOptions())
	if err != nil {
		fmt.Println(defaultHealthChecker.CheckAllFailureOrMessagef(
			"Error connecting to remote machine: '%s'", err,
		))
		return 1
	}

	if err := k.Dial(); err != nil {
		fmt.Println(defaultHealthChecker.CheckAllFailureOrMessagef(
			"Error connecting to remote machine: '%s'", err,
		))
		return 1
	}

	infos, err := getListOfMachines(k)
	if err != nil {
		fmt.Print(err)
		return 1
	}

	// allow for shortcuts when specifying name
	for _, info := range infos {
		if strings.HasPrefix(info.VMName, name) {
			name = info.VMName
		}
	}

	mountRequest := struct {
		Name       string `json:"name"`
		LocalPath  string `json:"localPath"`
		RemotePath string `json:"remotePath"`
		NoIgnore   bool   `json:"noIgnore"`
		NoPrefetch bool   `json:"noPrefetch"`
		NoWatch    bool   `json:"noWatch"`
	}{
		Name:       name,
		LocalPath:  localPath,
		NoIgnore:   noIgnore,
		NoPrefetch: noPrefetch,
		NoWatch:    noWatch,
	}

	// RemotePath is optional
	if remotePath != "" {
		mountRequest.RemotePath = remotePath
	}

	resp, err := k.Tell("remote.mountFolder", mountRequest)
	if err != nil && klientctlerrors.IsExistingMountErr(err) {
		util.MustConfirm("This folder is already mounted. Remount? [Y|n]")

		// unmount using mount path
		if err := unmount(k, "", localPath); err != nil {
			fmt.Printf("Error unmounting: %s\n", err)
			return 1
		}

		resp, err = k.Tell("remote.mountFolder", mountRequest)
		if err != nil {
			fmt.Println(defaultHealthChecker.CheckAllFailureOrMessagef(
				"Error mounting: %s\n", err,
			))
			return 1
		}
	}

	// catch errors other than klientctlerrors.IsExistingMountErr
	if err != nil {
		fmt.Println(defaultHealthChecker.CheckAllFailureOrMessagef(
			"Error mounting: %s\n", err,
		))
		return 1
	}

	// response can be nil even when there's no err
	if resp != nil {
		var warning string
		if err := resp.Unmarshal(&warning); err != nil {
			return 0
		}

		if len(warning) > 0 {
			fmt.Printf("Warning: %s\n", warning)
		}
	}

	if err := Lock(localPath, name); err != nil {
		fmt.Printf("Error locking: %s\n", err)
		return 1
	}

	fmt.Println("Mount success.")

	return 0
}

// askToCreate checks if the folder does not exist, and creates it
// if the user chooses to. If the user does *not* choose to create it,
// we return an IsNotExist error.
//
// TODO: Handle the case where a user types stuff in before being prompted,
// and then the prompt uses that. Ie, flush the input so that what we
// read is new input from the user. Not tested :)
func askToCreate(p string, r io.Reader, w io.Writer) error {
	_, err := os.Stat(p)

	// If we fail to stat the file, and it's *not* IsNotExist, we may be
	// having permission issues or some other related issue. Return
	// the error.
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// If there was no error stating the path, it already exists -
	// we can return, as there's nothing we need to do.
	if err == nil {
		return nil
	}

	fmt.Fprint(w,
		"The mount folder does not exist, would you like to create it? [Y/n]",
	)

	// To understand why we're creating a bReader here, please see
	// the docstring on YesNoConfirmWithDefault().
	bReader := bufio.NewReader(r)

	// Retry YesNo confirmation 3 times if needed
	var createFolder bool
	for i := 0; i < 3; i++ {
		createFolder, err = util.YesNoConfirmWithDefault(bReader, true)
		// If the user supplied an accepted value, stop trying
		if err == nil {
			break
		}
		// If err != nil, then the error did not provide an understood
		// response.
		fmt.Fprintln(w, "Invalid response, please type 'yes' or 'no'")
	}

	// If the retry loop exited with an error, the user failed to give
	// a meaningful response to the YesNo confirmation.
	if err != nil {
		return err
	}

	// The user chose not to create the folder. We cannot mount something that
	// doesn't exist - so we must fail here with an error.
	if !createFolder {
		return klientctlerrors.ErrUserCancelled
	}

	return os.Mkdir(p, 0655)
}
