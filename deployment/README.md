# Useful commands for deployment

### View systemd logs 
```shell
journalctl -f -u ecomapi.service
```
### Edit systemd file
``` 
sudo nano /etc/systemd/system/ecomapi.service
```
### Stop systemd service
```shell
systemctl stop ecomapi.service
```
### Check systemd status
```
shell systemctl status ecomapi.service
```

### Restart systemd service
```shell
systemctl restart ecomapi.service
```

### Connect database
```shell
sudo -u postgres psql

```
