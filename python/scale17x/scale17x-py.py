import datetime
import json
import sys


def main():
    now = datetime.datetime.now()
    if len(sys.argv) > 1:
        name = '{}'.format(sys.argv[1])
    else:
        name = 'Anonymous'
    output = {
        'msg': 'Hello {}! This is from inside the Python Binary'.format(name),
        'date': now.strftime("%Y-%m-%d %H:%M:%S")
    }
    print(json.dumps(output))


if __name__ == '__main__':
    main()
