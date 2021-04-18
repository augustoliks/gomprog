.. image:: https://github.com/augustoliks/gomprog/actions/workflows/go.yml/badge.svg?branch=main
    :target: https://github.com/augustoliks/gomprog/actions/workflows/go.yml

.. image:: https://codecov.io/gh/augustoliks/gomprog/branch/main/graph/badge.svg?token=aagkZMUNew
    :target: https://codecov.io/gh/augustoliks/gomprog

.. image:: https://goreportcard.com/badge/github.com/augustoliks/gomprog
    :target: https://goreportcard.com/report/github.com/augustoliks/gomprog
    :alt: Go Report Card

.. image:: https://pkg.go.dev/badge/github.com/augustoliks/gomprog.svg
    :target: https://pkg.go.dev/github.com/augustoliks/gomprog
    :alt: Go Reference

gomprog
=======

Rsyslog omprog plugin, created to forward logs to Redis instance. Plugin writted in Golang.

Dependencies
------------

============  ===============
Dependencies  Version
============  ===============
rsyslog       8.2010.0-1
============  ===============

How install
-----------

.. code-block:: shell

  go get github.com/augustoliks/gomprog

Rsyslog Config
--------------

.. code-block:: shell

  # Incoming log from udp syslog
  module(
    load="imudp"
  )

  # Import omprog, needded for fork gomprog binary
  module(
    load="omprog"
  )

  # Define log template formatter GELF, compatible with gomprog handler
  template(name="gelf" type="list") {
    constant(value="{\"_app_name\":\"")       property(name="app-name" caseconversion="lower")
    constant(value="\",\"host\":\"")          property(name="$myhostname")
    constant(value="\",\"short_message\":\"") property(name="msg" format="json")
    constant(value="\",\"timestamp\":")       property(name="timegenerated" dateformat="unixtimestamp")
    constant(value=",\"_group\":\"servers\"}\n")
  }

  # Import omprog, needded for fork gomprog binary
  input(
    ruleset="main"
    type="imudp" 
    port="10514"  
  )

  # Define sub-ruleset to delivery logs, from imudp to gomprog binary (rsyslog omprog)
  ruleset(name="main"){
    call gomprog
  #  call plugin-foo
  #  call ...
  #  call plugin-bar
  }

  # Define sub-ruleset to delivery gomprog binary (rsyslog omprog)
  ruleset(name="gomprog"){

    action(
      type="omprog"
      binary="gomprog-redis -n 127.0.0.1:6379 -p redis-password"
      template="gelf"
      action.resumeInterval="5" 
    #  output="/tmp/gomprog-debug.log"
    )

  }

Tests
-----

Download project

.. code-block:: shell

  git clone https://github.com/augustoliks/gomprog
  cd gomprog/
  
Provisioning redis instance

.. code-block:: shell

  cd tests/
  docker-compose up -d 

Configure Rsyslog

.. code-block:: shell

  cp packaging/rsyslog.d/30-gomprog-redis.conf /etc/rsyslog.d/30-gomprog-redis.conf
  systemctl restart rsyslog 

Log to upd using ``logger`` cli command

.. code-block:: shell

  logger -d -n 127.0.0.1 -p 10514 "tests"

Access Redis container, subscribe ``_app_name`` channel

.. code-block:: shell

  docker exec --user root -it redis bash                                                                                                           

  root@e7d850092677:/data# redis-cli 

  127.0.0.1:6379> SUBSCRIBE augustoliks
  Reading messages... (press Ctrl-C to quit)
  1) "subscribe"
  2) "augustoliks"
  3) (integer) 1
  1) "message"
  2) "augustoliks"
  3) "{\"host\":\"localhost\",\"short_message\":\"tests\",\"timestamp\":1618719119,\"_group\":\"servers\",\"_app_name\":\"augustoliks\"}"

.. note::

  **augustoliks** should be replace with Your Linux Username

Output expected

.. image:: ./docs/image/example.jpg
   :target: ./docs/image/example.jpg

References
----------

- https://petersouter.xyz/testing-and-mocking-stdin-in-golang/
- https://github.com/golang-standards/project-layout
