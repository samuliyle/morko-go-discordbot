# Commands

```
Command   Description

commands  Lists commands
google    Fetches a Google image
help      Describes the usage of the command.
ping      Pings
quote     Posts a random quote from the channel
remind    Sets a reminder
uptime    Bots uptime
```

## Secrets

Enter bot token to /config/secrets.json.

(optional) Enter Google and Database secrets to enable google and database commands.

## MySQL Database

Create database called 'discord' and import schema

```
mysql -u <username> -p
create database discord;
exit;
mysql -u <username> -p discord < schema.sql
```
