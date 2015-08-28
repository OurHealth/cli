# Catalyze CLI Version 2.0.0

## Upgrading from 1.X.X <a id="upgrading"></a>

Upgrading to the Catalyze CLI version 2.0.0 is easy! First, you need to uninstall the previous version of the CLI. This is most likely done through `pip uninstall catalyze`. Next you need to [download](https://github.com/catalyzeio/cli#downloads) the new version. Lastly, you need to [reassociate](#associate) all your environments.

## Supported Platforms & Architectures <a id="supported-platforms"></a>

Since version 2.0.0, the following platforms and architectures are supported by the Catalyze CLI.

| OS | Architecture |
|----|--------------|
| Darwin (Mac OS X) | 64-bit |
| Linux | 64-bit, arm |
| Windows | 64-bit |

# Global Scope <a id="global-scope"></a>

The CLI now supports the concept of scope. Previous to version 2.0.0, all commands had to be run within an associated local git repo. Now, the only time you need to be in a local git repo is when you associate a new environment. After the initial associated, CLI commands can be run from any directory. If you have more than one environment, the CLI uses this concept of scope to decide which environment you are using for the command. 

Let's say you have an environment that you associated in the directory `~/mysandbox-code` and another you associated in the directory `~/myprod-code`. These environments are named `mysandbox` and `myprod` respectively. When you are within either of those directories, the CLI knows that any command you run will be in the context of the given environment. Commands run in the `~/myprod-code` directory will be run against the `myprod` environment. Similarly for `~/mysandbox-code` and the `mysandbox` environment. What if you are outside those directories? You have three options.

First, you can tell the CLI which environment you want to use with the global option `-E` or `--env` (see [Global Options](#global-options)). Your command might start like this

```
catalyze -E myprod ...
```

This global option will even override the environment found in a local git repo. If you don't set the `-E` flag, and the CLI can't find an environment in your local git repo, the CLI then checks for a default environment. A default environment is used whenever you are outside of a git repo and an environment is not specified. A default environment can be specified using the [default](#default) command. You can find out which environment is the default by running the [associated](#associated) command.

Lastly, if no environment is specified, you're outside of a git repo, and no default environment is set, then the CLI simply takes the first environment you associated and prompts you to continue with this environment. This concept of scope will make it easier for Catalyze customers with multiple environments to use the CLI!

# Environment Aliases <a id="aliases"></a>

When you associate an environment from within a local git repo, you typically run the following command:

```
catalyze associate "My Health Tech Company Production"
```

Where `My Health Tech Company Production` is the name of your environment. However with the concept of [scope](#global-scope) and being able to specify which environment to use on a command by command basis with the `-E` global option, that is a lot to type! This is where environment aliases come in handy.

When you associate an environment and you want to pick a shorter name to reference the environment by, simply add a `-a` flag to the command. Let's try the command again calling it `prod` this time:

```
catalyze associate "My Health Tech Company Production" -a prod
```

Now when you run the [associated](#associated) command, you will see the alias as well as the actual environment name. 

When using aliases, there are a couple things to keep in mind. Aliases are only local and never leave your local machine. If you alias this environment `prod`, a coworker can alias the environment `healthtech-prod` with no ramifications. Second, after setting an alias, you will never reference the environment by its actual name with the CLI. You will always use the alias for flags, arguments, options, etc. 

To change or remove an alias, you must [disassociate](#disassociate) and then [reassociate](#associate) with a new alias.

# Bash Autocompletion <a id="autocompletion"></a>

One feature we've found helpful on \*Nix systems is autocompletion in bash. To enable this feature, head over to the github repo and download the `bash_autocomplete` file. If you use a Mac, you will need to install bash-completion with `brew install bash-completion` or `source` the `bash_autocomplete` file each time you start up terminal. Store this file locally in `/etc/bash_completion.d/` or (`/usr/local/etc/bash_completion.d/` on Mac). Completion will be available when you restart terminal. Now simply type `catalyze ` and hit tab twice to see the list of available commands. **Please note** that autocompletion only works one level deep. The CLI will not autocomplete or suggest completions when you type `catalyze db ` and then hit tab twice. It currently only works when you have just `catalyze ` typed into your terminal. This is a feature we are looking into expanding in the future.

Note: you may have to add `source /etc/bash_completion.d/completionfile` (`/usr/loca/etc/bsah_completion.d/completionfile`) in your `~/.bashrc` (`~/.bash_profile`) file.

# Global Options <a id="global-options"></a>

The following table outlines all global options available in the CLI. Global options are always set after the word `catalyze` and before any commands. Rather than setting these each time, you may also set an environment variable with the appropriate value which will automatically be used.

| Short Name | Long Name | Description | Environment Variable |
|------------|-----------|-------------|----------------------|
| -U | --username | Your catalyze username that you login to the Dashboard with | CATALYZE_USERNAME |
| -P | --password | Your catalyze password that you login to the Dashboard with | CATALYZE_PASSWORD |
| -E | --env | The local alias of the environment in which this command will be run. Read more about [environment aliases](#aliases) | CATALYZE_ENV |
| | --version | Prints out the CLI version | |

# Commands <a id="commands"></a>

This section lists all commands the CLI offers. Help text, along with a description, and a sample are given for each command.

## associate <a id="associate"></a>

```
Usage: catalyze associate ENV_NAME [SERVICE_NAME] [-a] [-r] [-d] 

Associates an environment

Arguments:
  ENV_NAME=""       The name of your environment
  SERVICE_NAME=""   The name of the primary code service to associate with this environment (i.e. 'app01')

Options:
  -a, --alias=""            A shorter name to reference your environment by for local commands
  -r, --remote="catalyze"   The name of the remote
  -d, --default=false       Specifies whether or not the associated environment will be the default
```

`associate` is the entry point of the cli. You need to associate an environment before you can run most other commands. Check out [scope](#global-scope) and [aliases](#aliases) for more info on the value of the alias and default options. Here is a sample command

```
catalyze associate My-Production-Environment -a prod -d
```

## associated <a id="associated"></a>

```
Usage: catalyze associated  

Lists all associated environments
```

`associated` outputs information about all previously associated environments on your local machine. The information that is printed out includes the alias, environment ID, actual environment name, service ID, the git repo directory, and whether or not it is the default environment. Here is a sample command

```
catalyze associated
```

## backup <a id="backup"></a>

The `backup` command gives access to backup and restores services for databases. The backup command can not be run directly but has four sub commands.

### create <a id="backup-create"></a>

```
Usage: catalyze backup create SERVICE_NAME [-s] 

Create a new backup

Arguments:
  SERVICE_NAME=""   The name of the database service to create a backup for (i.e. 'db01')

Options:
  -s, --skip-poll=false   Whether or not to wait for the backup to finish
```

`backup create` creates a new backup for the given database service. The backup is started and unless `-s` is specified, the CLI will poll every 2 seconds until it finishes. Regardless of a successful backup or not, the logs for the backup will be printed to the console when the backup is finished. Here is a sample command

```
catalyze backup create db01
```

### download <a id="backup-download"></a>

```
Usage: catalyze backup download SERVICE_NAME BACKUP_ID FILEPATH [-f] 

Download a previously created backup

Arguments:
  SERVICE_NAME=""   The name of the database service which was backed up (i.e. 'db01')
  BACKUP_ID=""      The ID of the backup to download (found from "catalyze backup list")
  FILEPATH=""       The location to save the downloaded backup to. This location must NOT already exist unless -f is specified

Options:
  -f, --force=false   If a file previously exists at "filepath", overwrite it and download the backup
```

`backup download` downloads a previously created backup to your local hard drive. Be careful using this command is it could download PHI. Be sure that all hard drive encryption and necessary precautions have been taken before performing a download. The ID of the backup is found by first running the [backup list](#backup-list) command. Here is a sample command

```
catalyze backup download db01 cd2b4bce-2727-42d1-89e0-027bf3f1a203 ./db.sql
```

This assumes you are download a MySQL or PostgreSQL backup which takes the `.sql` file format. If you are downloading a mongo backup, the command might look like this

```
catalyze backup download db01 cd2b4bce-2727-42d1-89e0-027bf3f1a203 ./db.tar.gz
```

### list <a id="backup-list"></a>

```
Usage: catalyze backup list SERVICE_NAME [-p] [-n] 

List created backups

Arguments:
  SERVICE_NAME=""   The name of the database service to list backups for (i.e. 'db01')

Options:
  -p, --page=1         The page to view
  -n, --page-size=10   The number of items to show per page
```

`backup list` lists all previously created backups. After listing backups you can copy the backup ID and use it to download that backup or restore your database from that backup. Here is a sample command

```
catalyze backup list db01
```

### restore <a id="backup-restore"></a>

```
Usage: catalyze backup restore SERVICE_NAME BACKUP_ID [-s] 

Restore from a previously created backup

Arguments:
  SERVICE_NAME=""   The name of the database service to restore (i.e. 'db01')
  BACKUP_ID=""      The ID of the backup to restore (found from `catalyze backup list`)

Options:
  -s, --skip-poll=false   Whether or not to wait for the restore to finish
```

`backup restore` restores your database service to a previous point in time. **Be careful** using this command as your database will be wiped. The backup ID can be found by running the [backup list](#backup-list) command. Regardless of a successful backup or not, the logs for the backup will be printed to the console when the backup is finished. Here is a sample command

```
catalyze backup restore db01 cd2b4bce-2727-42d1-89e0-027bf3f1a203
```
## console <a id="console"></a>

```
Usage: catalyze console SERVICE_NAME [COMMAND] 

Open a secure console to a service

Arguments:
  SERVICE_NAME=""   The name of the service to open up a console for
  COMMAND=""        An optional command to run when the console becomes available
```

`console` gives you direct access to your database service or application shell. For example, if you open up a console to a postgres database, you will be given access to a psql prompt. You can also open up a mysql prompt, mongo cli prompt, rails console, django shell, and much more. When accessing a database service, the `COMMAND` argument is not needed because the appropriate prompt will be given to you. If you are connecting to an application service the `COMMAND` argument is required. Here are some sample commands

```
catalyze console db01
catalyze console app01 "bundle exec rails console"
```

## dashboard <a id="dashboard"></a>

```
Usage: catalyze dashboard  

Open the Catalyze Dashboard in your default browser
```

`dashboard` simply opens up the Catalyze Dashboard homepage in your default web browser. Here is a sample command

```
catalyze dashboard
```

## db <a id="db"></a>

The `db` command gives access to import and export services for databases. The db command can not be run directly but has two sub commands.

### import <a id="db-import"></a>

```
Usage: catalyze db import DATABASE_NAME FILEPATH [-d [-c]] 

Import data to a database

Arguments:
  DATABASE_NAME=""   The name of the database to import data to (i.e. 'db01')
  FILEPATH=""        The location of the file to import to the database

Options:
  -c, --mongo-collection=""   If importing into a mongo service, the name of the collection to import into
  -d, --mongo-database=""     If importing into a mongo service, the name of the database to import into
  -w, --wipe-first=false      Whether or not to wipe the database before processing the import file
```

`import` allows you to inject new data into your database service. For example, if you wrote a simple SQL file

```
CREATE TABLE mytable (
id TEXT PRIMARY KEY,
val TEXT
);

INSERT INTO mytable (id, val) values ('1', 'test');
```

and stored it at `./db.sql` you could import this into your database service. If you are initializing your database, you will want to use the `-w` option as this will first remove all data in your database and then perform the import. When import data into mongo, you may specify the database and collection to import into using the `-d` and `-c` flags respectively. Regardless of a successful import or not, the logs for the import will be printed to the console when the import is finished. Here is a sample command

```
catalyze db import db01 ./db.sql
```

### export <a id="db-export"></a>

```
Usage: catalyze db export DATABASE_NAME FILEPATH [-f] 

Export data from a database

Arguments:
  DATABASE_NAME=""   The name of the database to export data from (i.e. 'db01')
  FILEPATH=""        The location to save the exported data. This location must NOT already exist unless -f is specified

Options:
  -f, --force=false   If a file previously exists at `filepath`, overwrite it and export data
```

`export` is a simple wrapper around the `backup create` and `backup download` command. When you request an export, a backup is created that will be added to the list of backups shown when you perform the [backup list](#backup-list) command. Next, that backup is immediately downloaded. Regardless of a successful export or not, the logs for the export will be printed to the console when the export is finished. Here is a sample command

```
catalyze db export db01 ./dbexport.sql
```

## default <a id="default"></a>

```
Usage: catalyze default ENV_ALIAS 

Set the default associated environment

Arguments:
  ENV_ALIAS=""   The alias of an already associated environment to set as the default
```

`default` sets the default environment for all commands without a specified environment and run outside of a git repo. See [scope](#global-scope) for more information on scope and default environments. When setting a default environment, you must give the alias of the environment if one was set when it was associated and not the real environment name. Here is a sample command

```
catalyze default prod
```

## disassociate <a id="disassociate"></a>

```
Usage: catalyze disassociate  

Remove the association with an environment
```

`disassociate` must be run from within an associated git repo. Disassociate removes the environment from your list of associated environments but **does not** remove the catalyze git remote on the git repo. Here is a sample command

```
catalyze disassociate
```

## environments <a id="environments"></a>

```
Usage: catalyze environments  

List all environments you have access to
```

`environments` lists all environments that you are granted access to. These environments include those you created and those that other Catalyze customers have added you to. Here is a sample command

```
catalyze environments
```

## logs <a id="logs"></a>

```
Usage: catalyze logs [QUERY] [-f] [-h] 

Show the logs in your terminal streamed from your logging dashboard

Arguments:
  QUERY="app*"   The query to send to your logging dashboard's elastic search (regex is supported)

Options:
  -f, --follow=false   Tail/follow the logs
  -h, --hours=4        The number of hours to retrieve logs
```

`logs` prints out your application logs directly from your logging Dashboard. If you do not see your logs, try adjusting the number of hours of logs that are retrieved with the `-h` option. You can also follow the logs with the `-f` option. When using `-f` all logs will be printed to the console within the given time frame as well as any new logs that are sent to the logging Dashboard for the duration of the command. When using the `-f` option, hit ctrl-c to stop. Here is a sample command

```
catalyze logs -f -h 24
```

## logout <a id="logout"></a>

```
Usage: catalyze logout  

Clear the stored user information from your local machine
```

When using the CLI, your username and password are **never** stored in any file on your filesystem. However, in order to not type in your username and password each and every command, a session token is stored in the CLI's configuration file and used until it expires. `logout` removes this session token from the configuration file. Here is a sample command

```
catalyze logout
```

## metrics <a id="metrics"></a>

```
Usage: catalyze metrics [SERVICE_NAME] [(--json | --csv | --spark)] [--stream] [-m] 

Print service and environment metrics in your local time zone

Arguments:
  SERVICE_NAME=""   The name of the service to print metrics for

Options:
  --json=false     Output the data as json
  --csv=false      Output the data as csv
  --spark=false    Output the data using spark lines
  --stream=false   Repeat calls once per minute until this process is interrupted.
  -m, --mins=1     How many minutes worth of logs to retrieve.
```

`metrics` prints out various metrics for your environment or individual service. Metrics included are CPU metrics, Memory metrics, Disk I/O metrics, and Network metrics. You can print out metrics in csv, json, plain text, or spark lines format. If you want plain text format, simply omit the json, csv, and spark options. You can only stream metrics using plain text or spark lines formats. To print out metrics for every service in your environment, omit the `SERVICE_NAME` argument. Otherwise you may choose a service, such as an app service, to retrieve metrics for. Here are some sample commands

```
catalyze metrics
catalyze metrics app01 --stream
catalyze metrics --json
catalyze metrics db01 --csv -m 60
```

## rake <a id="rake"></a>

```
Usage: catalyze rake TASK_NAME 

Execute a rake task

Arguments:
  TASK_NAME=""   The name of the rake task to run
```

`rake` executes a rake task by its name asynchronously. Once executed, the output of the task can be seen through your logging Dashboard or using the [logs](#logs) command. Here is a sample command

```
catalyze rake db:migrate
```

## redeploy <a id="redeploy"></a>

```
Usage: catalyze redeploy  

Redeploy a service without having to do a git push
```

`redeploy` restarts your code service without having to perform a code push. Typically when you want to update your code service you make a code change, git commit, then git push catalyze master. After the build finishes and a couple minutes later your code service will be redeployed. With the redeploy command, you skip the git push and the build. The code service that is redeployed is the one that was associated when the [associate](#associate) command was run. You can see which code service ID this is with the [associated](#associated) command. Here is a sample command

```
catalyze redeploy
```

## ssl <a id="ssl"></a>

### verify <a id="ssl-verify"></a>

```
Usage: catalyze ssl verify CHAIN PRIVATE_KEY HOSTNAME [-s] 

Verify whether a certificate chain is complete and if it matches the given private key

Arguments:
  CHAIN=""         The path to your full certificate chain in PEM format
  PRIVATE_KEY=""   The path to your private key in PEM format
  HOSTNAME=""      The hostname that should match your certificate (i.e. "*.catalyze.io")

Options:
  -s, --self-signed=false   Whether or not the certificate is self signed. If set, chain verification is skipped
```

`ssl verify` will tell you if your SSL certificate and private key are properly formatted for use with the Catalyze PaaS. Before uploading a certificate to Catalyze you should verify it creates a full chain and matches the given private key with this command. Both your chain and private key should be **unencrypted** and in **pem** format. The private key is the only key in the key file. However, for the chain, you should include your SSL certificate, intermediate certificates, and root certificate in the following order and format.

```
-----BEGIN CERTIFICATE-----
<Your SSL certificate here>
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
<One or more intermediate certificates here>
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
<Root CA here>
-----END CERTIFICATE-----
```

This command also requires you to specify the hostname that you are using the SSL certificate for in order to verify that the hostname matches what is in the chain. If it is a wildcard certificate, your hostname would be in the following format: `*.catalyze.io`. This command will verify a complete chain can be made from your certificate down through the intermediate certificates all the way to a root certificate that you have given or one found in your system. 

You can also use this command to verify self-signed certificates match a given private key. To do so, add the `-s` option which will skip verifying the certificate to root chain and just tell you if your certificate matches your private key. Here is a sample command

```
catalyze ssl verify ./catalyze.crt ./catalyze.key *.catalyze.io
```

## status <a id="status"></a>

```
Usage: catalyze status  

Get quick readout of the current status of your associated environment and all of its services
```

`status` will give a quick readout of your environment's health. This includes your environment name, environment ID, and for each service the name, size, build status, deploy status, and service ID. Here is a sample command

```
catalyze status
```

## support-ids <a id="support-ids"></a>

```
Usage: catalyze support-ids  

Print out various IDs related to your associated environment to be used when contacting Catalyze support
```

`support-ids` is helpful when contacting Catalyze support by sending an email to support@catalyze.io. If you are having an issue with a CLI command or anything with your environment, it is helpful to run this command and copy the output into the initial correspondence with a Catalyze engineer. This will help Catalyze identify the environment faster and help come to resolution faster. Here is a sample command

```
catalyze support-ids
```

## users <a id="users"></a>

The `users` command allows you to manage who has access to your environment. The users command can not be run directly but has three sub commands.

### add <a id="users-add"></a>

```
Usage: catalyze users add USER_ID 

Grant access to the associated environment for the given user

Arguments:
  USER_ID=""   The Users ID to give access to the associated environment
```

`users add` grants an existing Catalyze Dashboard user access to your environment. To give them access, request that they first run the [whoami](#whoami) command and send you their users ID. Here is a sample command

```
catalyze users add 774bf982-fc4a-428b-a048-c38cffb7d0ab
```

### list <a id="users-list"></a>

```
Usage: catalyze users list  

List all users who have access to the associated environment
```

`users list` shows every user that has access to your environment. Only the users ID of each user is printed out. Here is a sample command

```
catalyze users list
```

### rm <a id="users-rm"></a>

```
Usage: catalyze users rm USER_ID 

Revoke access to the associated environment for the given user

Arguments:
  USER_ID=""   The Users ID to revoke access from for the associated environment
```

`users rm` revokes a users access to your environment. This is the opposite of the [users add](#users-add) command. Here is a sample command

```
catalyze users rm 774bf982-fc4a-428b-a048-c38cffb7d0ab
```

## vars <a id="vars"></a>

The `vars` command allows you to manage environment variables for your code services. The vars command can not be run directly but has three sub commands.

### list <a id="vars-list"></a>

```
Usage: catalyze vars list  

List all environment variables
```

`vars list` prints out all known environment variables for the associated code service. Here is a sample command

```
catalyze vars list
```

### set <a id="vars-set"></a>

```
Usage: catalyze vars set -v... 

Set one or more new environment variables or update the values of existing ones

Options:
  -v, --variable    The env variable to set or update in the form "<key>=<value>"
```

`vars set` allows you to add a new environment variable or update the value of an existing environment variable on your code service. You can set/update 1 or more environment variables at a time with this command by repeating the `-v` option multiple times. Once new environment variables are added or values updated, a [redeploy](#redeploy) is required for your code service to have access to the changes. The environment variables must be of the form `<key>=<value>`. Here is a sample command

```
catalyze vars set -v AWS_ACCESS_KEY_ID=1234 -v AWS_SECRET_ACCESS_KEY=5678
```

### unset <a id="vars-unset"></a>

```
Usage: catalyze vars unset VARIABLE 

Unset (delete) an existing environment variable

Arguments:
  VARIABLE=""   The name of the environment variable to unset
```

`vars unset` removes an environment variables from your associated code service. Only the environment variable name is required to unset. Once environment variables are unset, a [redeploy](#redeploy) is required for your code service to have access to the changes. Here is a sample command

```
catalyze vars unset AWS_ACCESS_KEY_ID
```

## version <a id="version"></a>

```
Usage: catalyze version  

Output the version and quit
```

`version` prints out the current CLI version. Here is a sample command

```
catalyze version
```

## whoami <a id="whoami"></a>

```
Usage: catalyze whoami  

Retrieve your user ID
```

`whoami` prints out the currently logged in user's users ID. This is used with the [users add](#users-add) and [users rm](#users-rm) commands as well as with Catalyze support. Here is a sample command

```
catalyze whoami
```

## worker <a id="worker"></a>

```
Usage: catalyze worker TARGET 

Start a background worker

Arguments:
  TARGET=""    The name of the Procfile target to invoke as a worker
```

`worker` starts a background worker asynchronously. The `TARGET` argument must be specified in your `Procfile`. Once the worker is started, any output can be found in your logging Dashboard or using the [logs](#logs) command. Here is a sample command

```
catalyze worker Scrape
```