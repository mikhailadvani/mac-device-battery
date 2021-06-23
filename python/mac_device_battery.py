import argparse
import subprocess
import json

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--device-type', default='SPBluetoothDataType', dest='device_type', help='Device type whose battery should be checked. Checks all SPBluetoothDataType by default')
    args = parser.parse_args()

    device_info = subprocess.check_output(["system_profiler", "-json", args.device_type], stderr=subprocess.DEVNULL)
    device_info_dict = json.loads(device_info)

    data_type = device_info_dict[args.device_type][0]
    for device in data_type['device_title']:
        for device_name, device_spec in device.items():
            if device_spec.get('device_batteryPercent') is not None:
                print(device_name, device_spec['device_batteryPercent'])
    # print(data_type)

if __name__ == '__main__':
    main()
