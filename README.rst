gomprog
=======

Rsyslog omprog plugin, created to forward logs to Redis instance. Plugin writted in Golang.

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
