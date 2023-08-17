
### REQUIREMENTS
This program will only work on UNIX systems where the battery 'capacity' and 'status' files can be found at '/sys/class/power_supply/BAT0/'
It simply calls [twmn](https://github.com/sboli/twmn) to send notifications when the battery falls below certain thresholds.

### SET UP
Make sure that [twmn](https://github.com/sboli/twmn) is installed on your system and the twmn daemon is running.  It must be launched with your DE or WM.  It cannot be set up as a systemd service.  
Install lowbatteryalert using 'go install'.  Then set it up as a systemd service.
