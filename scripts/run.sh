#!/bin/bash
export $(cat ../configs/example.env | xargs)

echo DONE!