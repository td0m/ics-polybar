# ICS Polybar

## Screenshots

TODO: add screenshots

## Usage

### Customise and Build your version

```
git clone https://github.com/d0minikt/ics-polybar
cd ics-polybar
go build
```

### Move the binary

```
mkdir -p ~/.config/polybar/scripts
cp ics-polybar ~/.config/polybar/scripts/ics-polybar
```

### Add `calendar.ics`

```
cp path/to/calendar.ics ~/.config/polybar/scripts/calendar.ics
```

### Add Polybar Script Module

```
[module/calendar]
type = custom/script
exec = ~/.config/polybar/scripts/ics-polybar
interval = 60
```

