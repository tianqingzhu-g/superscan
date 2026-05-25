#!/usr/bin/env node
const fs = require('fs');
const path = process.argv[2] || 'package-lock.json';
if (!fs.existsSync(path)) process.exit(0);
const data = JSON.parse(fs.readFileSync(path, 'utf8'));
const result = new Map();

function processDeps(deps) {
  if (!deps) return;
  for (const [name, meta] of Object.entries(deps)) {
    if (!meta) continue;
    const ver = meta.version || meta.resolved || 'unknown';
    if (ver === 'unknown') continue;
    const key = name.toLowerCase();
    if (!result.has(key)) result.set(key, ver);
    if (meta.dependencies) processDeps(meta.dependencies);
  }
}

if (data.dependencies) processDeps(data.dependencies);
if (data.packages) {
  for (const [pkgPath, meta] of Object.entries(data.packages)) {
    if (pkgPath === '') continue;
    const name = meta.name || (pkgPath.split('node_modules/').pop());
    if (!name) continue;
    const ver = meta.version || 'unknown';
    if (ver === 'unknown') continue;
    result.set(name.toLowerCase(), ver);
  }
}

for (const [name, version] of result.entries()) {
  console.log(JSON.stringify({ ecosystem: 'npm', name, version }));
}