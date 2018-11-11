Example Curl Request to Add a Pokemon:
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Charizard","item":"Charcoal","ability":"Blaze","level":50,"hitpointsev":12,"attackev":16,"defenseev":20,"specialattackev":24,"specialdefenseev":28,"speedev":32,"nature":"Timid","hitpointsiv":20,"attackiv":0,"defenseiv":18,"specialattackiv":31,"specialdefenseiv":28,"speediv":31,"moves":["Flamethrower","Heat Wave","Tailwind","Protect"]}' \
  http://localhost:8000/api/pokemon/add