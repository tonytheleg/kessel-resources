#!/usr/bin/env python

import yaml, argparse
from jinja2 import Template

parser = argparse.ArgumentParser(description='A simple argparse example.')
parser.add_argument('template_file', help='path to the template file to render')
parser.add_argument('config_file', help='path to the config file containing variables to inject into template')
args = parser.parse_args()

with open(args.config_file, 'r') as file:
    data = yaml.load(file, Loader=yaml.SafeLoader)

with open(args.template_file, 'r') as file:
    template_data = file.read()

template = Template(template_data)
rendered = template.render(data)

print(rendered)
