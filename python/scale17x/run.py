import datetime
import json


now = datetime.datetime.now()
output = {
    'msg': 'Hello from inside the Python Binary',
    'date': now.strftime("%Y-%m-%d %H:%M:%S")
}
print(json.dumps(output))
