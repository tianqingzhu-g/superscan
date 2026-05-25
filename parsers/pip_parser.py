#!/usr/bin/env python3
import json, sys, os

def parse_pipfile_lock(path):
    with open(path, 'r', encoding='utf-8') as f:
        data = json.load(f)
    result = []
    def extract(sec):
        for name, meta in (data.get(sec) or {}).items():
            if isinstance(meta, str):
                ver = meta.strip().lstrip("=")
                if any(op in ver for op in ('>','<','!','~',';')):
                    ver = 'unknown'
            elif isinstance(meta, dict):
                if 'version' in meta:
                    ver = meta['version'].lstrip("=")
                elif 'ref' in meta:
                    ver = meta['ref']
                elif 'path' in meta:
                    ver = meta['path']
                else:
                    ver = 'unknown'
            else:
                ver = 'unknown'
            if ver != 'unknown':
                result.append((name.replace('_','-').lower(), ver))
    extract('default')
    extract('develop')
    seen = {}
    for n, v in result:
        if n not in seen:
            seen[n] = v
    return [(k, v) for k, v in seen.items()]

if __name__ == '__main__':
    p = sys.argv[1] if len(sys.argv) > 1 else 'Pipfile.lock'
    if not os.path.exists(p):
        sys.exit(0)
    for name, ver in parse_pipfile_lock(p):
        print(json.dumps({"ecosystem": "PyPI", "name": name, "version": ver}))