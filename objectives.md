# Root cleaner (can be applied to any directory)

The purpose of this program is to move (selected) files from the home directory to ~/Desktop (or any
other directory according to the XDG protocol) to the desktop directory.
Should there already be an entry (directory or file) in the Desktop directory, the program should
either ask what to with that file (override, save both, preserve the original file) or apply the
predefined settings (from the config file or specified by an environment variable).
