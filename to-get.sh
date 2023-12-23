#! /bin/bash
#Using jq to filter data extracted from a json file using multiple parameters.


curl https://learn.zone01kisumu.ke/assets/superhero/all.json | jq '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'