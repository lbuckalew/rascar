# rascar

## Notes
1. Install Raspbian Jessie Lite
2. Add file "ssh" to filesystem to enable ssh on startup
3. sudo apt-get update
4. sudo apt-get install bluez pulseaudio-module-bluetooth python-gobject python-gobject-2
5. sudo usermod -a -G lp pi ; to be able to view bt sources and sinks etc
6. /etc/pulse/daemon.conf declare resample method
7. /etc/default/pulseaudio >    PULSEAUDIO_SYSTEM_START=1 DISALLOW_MODULE_LOADING=0
sudo adduser pi pulse-access
sudo nano /etc/pulse/client.conf > autospawn = no
sudo nano /etc/pulse/daemon.conf >    allow-module-loading = yes load-default-script-file = yes default-script-file = /etc/pulse/default.pa
sudo nano /etc/dbus-1/system.d/pulseaudio-system.conf > 
    <policy user="pulse">
          <allow own="org.pulseaudio.Server"/>
          <allow send_destination="org.bluez"/>
          <allow send_interface="org.bluez.Manager"/>
    </policy> EOF
8. pact list sources|sinks short
9. pactl load-module module-loopback source=bluez_source.XX_XX_XX_XX_XX_XX sink=alsa_output.platform-bcm2835_AUD0.0.analog-stereo
10. amixer cset numid=3 1 to set output to headphone jack
11. amixer set Master 100%; pacmd set-sink-volume 0 65537;
