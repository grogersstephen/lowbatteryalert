### LOW BATTERY ALERT

# DEPENDENCIES
This program will only work on UNIX systems where the battery 'capacity' and 'status' files can be found at '/sys/class/power_supply/BAT0/'
It simply calls [TWMN](https://github.com/sboli/twmn) to send notifications when the battery falls below certain thresholds.

# SET UP
Install using 'go install'.  Then set up as a systemd service.
