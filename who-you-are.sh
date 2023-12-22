#! /bin/bash
#This code uses the id given to fetch out the name corresponding to the id.

curl https://learn.zone01kisumu.ke/assets/superhero/all.json | jq '.[] | select(.id == 70) |.name' 
