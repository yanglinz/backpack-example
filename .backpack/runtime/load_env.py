import argparse
import json
import os
import shlex


def parse_json(json_path):
    with open(json_path) as f:
        data = json.load(f)
        return data


def output_bash_commands(json_path):
    for k, v in parse_json(json_path).items():
        value = shlex.quote(str(v))
        export = f"export {k}={value}"
        print(export)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("json_path", type=str)
    args = parser.parse_args()
    output_bash_commands(args.json_path)
