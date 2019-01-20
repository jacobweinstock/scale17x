import datetime
import json
import sys


now = datetime.datetime.now()
if len(sys.argv) > 1:
    name = '{}'.format(sys.argv[1])
else:
    name = 'Anonymous'
output = {
    'msg': 'Hello {}! Coming to you from inside the Python Binary'.format(name),
    'date': now.strftime("%Y-%m-%d %H:%M:%S")
}
print(json.dumps(output))
