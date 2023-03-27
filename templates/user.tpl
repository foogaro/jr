{{$NAME := name}}{{$SURNAME := surname}}{{$COMPANY := company}}{
  "guid": "{{uuid}}",
  "isActive": {{bool}},
  "balance": {{floating 1000 4000}},
  "picture": "http://placehold.it/32x32",
  "age": {{integer 20 60}},
  "eyeColor": "{{randoms "blue|brown|green"}}",
  "name": "{{$NAME}} {{$SURNAME}}",
  "company": "{{$COMPANY}}",
  "email": "{{lower $NAME}}.{{lower $SURNAME}}@{{firstword (lower $COMPANY)}}.com",
  "alt_email": "{{first (lower $NAME)}}.{{lower $SURNAME}}@{{squeeze (lower $COMPANY)}}.com",
  "about": "{{lorem 20}}",
  "latitude": {{floating -90.000001 90}},
  "longitude": {{floating -180.000001 180}}
}
