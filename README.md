# Footcon

Footcon is a USB device that send keyboard-like keys to a PC with a foot pedal.

<img src="images/footcon.png" width="70%" />

## Hardware

The hardware composition is based on References 1. (Japanese website).  

Composition detail is below:

* Housing
    * TAKACHI YM-100
* Substrate
    * [footcon](https://github.com/akif999/footcon/tree/main/kicad/footcon)
    * [gerber](https://github.com/akif999/footcon/tree/main/gerber)
        * A board can be made using the files in this folder
* Foot pedal
    * [YAMAHA FC-5](https://jp.yamaha.com/products/music_production/accessories/fc5/index.html)
* Microcontroller board
    * [Seeed Studio XIAO SAMD21](https://lab.seeed.co.jp/entry/2022/08/19/120000)


## Software

The software consists of the following two items:

* Firmware
    * Software to detect pedal input and send key to PC
        * Write to XIAO
* Config GUI
    * Software for setting up Footcon
        * Works standalone on PC

### Firmware

Build and use the tinygo project under [firmware](https://github.com/akif999/footcon/tree/main/firmware).  

```
$ cd firmware
$ tinygo flash -target=xiao
```

* Requirement
    * Tinygo 0.27 or later

### Config GUI

Build [config_gui](https://github.com/akif999/footcon/tree/main/config_gui) and below to create an executable and use it.

```
$ cd config_gui
$ go build -o config_gui
$ ./config_gui
```

* Requirement
    * Go 1.19 or later

## References

1. https://qiita.com/carcon999/items/9b57537cadbfd3c13cf0

## License

This project is licensed under the MIT license.

## Author

Akifumi Kitabatake (a.k.a akif or akif999)
