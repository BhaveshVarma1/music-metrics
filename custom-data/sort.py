# NOTE: This file relies on another Python file called owners.py
# owners.py must contain a dictionary called 'device_owners' that maps device names to owner names
# The owner names will become the names of the new json files that will be output by this file, so make them unique
# owners.py also contains the names of the old and new directories containing the json files, to sort and to output
#
# It also contains a dictionary called 'custom_rules' that maps a device name to an array of tuples.
# Each tuple has 1. a timestamp formatted like '2016-07-16T20:18:56Z' and a device owner's name. This program will
# partition the listens from that device by timestamps listed in this array (so the first timestamp should be
# something from before any listens were recorded, like 1900-01-01T00:00:00Z). The timestamps must be in order.
#
# Finally, there can be no owners named 'other' (see the get_owner method below).

import json
import os
from owners import device_owners, old_dir, new_dir, custom_rules

def print_specific_device(sorted_data, device_name):
  specific = []
  print(len(sorted_data))
  for owner in sorted_data:
    print(owner)
    for obj in sorted_data[owner]:
      if obj['platform'] == device_name:
        specific.append(obj)

  print(len(specific))
  counter = 0
  for obj in specific:
    if counter > 1500:
      break
    if obj['master_metadata_album_artist_name'] is not None:
      counter += 1
      print(f"{obj['master_metadata_track_name']} by {obj['master_metadata_album_artist_name']} on {obj['ts']}")

def print_specific_owner(sorted_data, owner):
  counter = 0
  for obj in sorted_data[owner]:
    if counter > 1500:
      break
    if obj['master_metadata_album_artist_name'] is not None:
      counter += 1
      print(f"{obj['master_metadata_track_name']} by {obj['master_metadata_album_artist_name']} on {obj['ts']}")

def print_devices(dir):

  devices = {}
  for file in os.listdir(dir):
    if not file.endswith('.json'):
      continue

    with open(os.path.join(dir, file), 'r') as json_file:
      data = json.load(json_file)

      for obj in data:
        device = obj['platform']

        if device in devices:
          devices[device] += 1
        else:
          devices[device] = 0

  print(list(devices.keys()))

def get_owner(all_owners, device, timestamp, rules):
  owner = 'other'
  if device in all_owners:
    owner = all_owners[device]
  elif device in custom_rules:
    for tup in custom_rules[device]:
      if timestamp > tup[0]:
        owner = tup[1]

  return owner

# if __name__ == '__main__':

sorted = {}

for file in os.listdir(old_dir):
  if not file.endswith('.json'):
    continue

  with open(os.path.join(old_dir, file), 'r') as json_file:

    # Iterate through each object in this json file
    for obj in json.load(json_file):
      owner = get_owner(device_owners, obj['platform'], obj['ts'], custom_rules)

      if owner in sorted:
        sorted[owner].append(obj)
      else:
        sorted[owner] = [obj]

# Everything has been sorted, write each to a json file
object_limit = 16000

for owner in sorted:
  data = sorted[owner]
  num_partitions = 1
  if len(data) > object_limit:
    num_partitions = (len(data) // object_limit) + 1

  cursor = 0
  for i in range(num_partitions):
    with open(os.path.join(new_dir, owner + f"_{i+1}.json"), 'w') as json_file:
      json.dump(sorted[owner][cursor:cursor + object_limit], json_file)
    cursor += object_limit
