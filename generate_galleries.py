#!/usr/bin/env python

import json
import os
import subprocess
import shlex


def generate_thumbnail(member, image_name):
    image_path = 'galleries/%s/images/%s' % (member, image_name)
    thumbnail_path = 'galleries/%s/thumbs/%s' % (member, image_name)
    command = ("convert -thumbnail '180x180^' -gravity center "
               "-extent 180x180 %s %s" % (image_path, thumbnail_path))

    status = subprocess.call(shlex.split(command))
    if status != 0:
        print('Failed to generate thumbnail for %s' % image_path)


def generate_thumbnails(member, image_names):
    if len(image_names) == 0:
        return

    thumbs_dir = "galleries/%s/thumbs" % member

    if not os.path.isdir(thumbs_dir):
        os.makedirs(thumbs_dir)

    existing_thumbnails = os.listdir(thumbs_dir)
    if len(existing_thumbnails) == len(image_names):
        return
    existing_thumbnails_map = {name: True for name in existing_thumbnails}

    print("Generating thumbnails for %s" % member)
    for image in image_names:
        if image not in existing_thumbnails_map:
            generate_thumbnail(member, image)


def output_gallery_json(images):
    for user, image_list in images.iteritems():
        img_objs = [{"filename": img} for img in image_list]
        to_write = {"images": img_objs}

        with open("galleries/%s/images.json" % user, "w") as f:
            json.dump(to_write, f)


print("Starting generation")

if not os.path.isdir('galleries'):
    print('Cannot find galleries folder.')
    print('Must be run from within the website directory.')
    exit(1)

images = {}

for member in os.listdir('galleries'):
    print('Looking for images for %s' % member)

    if not os.path.isdir('galleries/%s/images' % member):
        print('Cannot find images for %s. Skipping...' % member)
        continue

    images[member] = os.listdir('galleries/%s/images' % member)
    generate_thumbnails(member, images[member])

output_gallery_json(images)
