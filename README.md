# twtui

A simple TUI for the Seqera Platform.

## Usage

`twtui` automatically loads `TOWER_API_ENDPOINT` and `TOWER_ACCESS_TOKEN` from your
environment. Set these values, then simply run `go run .`

* The TUI supports basic navigation with h/j/k/l and arrow keys.
* Press [enter] to open the selected row in a new table view.
* Press [esc] to go back to the previous view.
* For paginated views, press [tab] and [shift]+[tab] to navigate between pages.
