from __future__ import print_function
import frida
import sys
import getopt


def usage():
    print("fuck off i haven't written a usage message yet")

try:
    opts, args = getopt.getopt(sys.argv[1:], "hs:", ["help", "script"])
except getopt.GetoptError as err:
    print(str(err))  # will print something like "option -a not recognized"
    usage()
    sys.exit(2)
script = None
for o, a in opts:
    if o in ("-h", "--help"):
        usage()
        sys.exit()
    elif o in ("-s", "--script"):
        script = a
    else:
        assert False, "unhandled option"

print("sending script: " + script)

with open(script, "r") as s:
    contents = s.read()

    print("setting up session....")
    device = frida.get_device_manager().enumerate_devices()[-1]
    session = device.attach("Gadget")
    print("done")
    print("send up script...")
    script = session.create_script(contents)
    print("done")
    def on_message(message, data):
        print(message)
    script.on('message', on_message)
    script.load()
    sys.stdin.read()
