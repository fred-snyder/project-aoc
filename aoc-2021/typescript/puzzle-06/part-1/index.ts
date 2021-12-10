import { readFileSync } from 'fs';

const PATH = '../puzzle-input/input-p6p1.txt';

const file = readFileSync(PATH, 'utf-8');

console.log(file);
