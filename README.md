Problem Overview
Each scooter has a unique id, and its current location (lat, lng). Users should be able to view scooters, start a scooter reservation, end their reservation, and pay for the distance they traveled. Build REST APIs for the following and share the Git repository with us. You can populate your database with any dummy data you want. You can write the code in Python/Django, Ruby/Rails, JS/Express or another web framework, but we have a preference for Python/Django.
Requirements
Search for an address and find nearby available scooters. (input: lat, lng, radius in meters. Output - list of parking spots within the radius).
Reserve a scooter.
Bonus
Automated tests
Ending reservations
Any kind of mock payments
Proposals on how to improve the APIs
Sample API requests/responses:
1.) GET /api/v1/scooters/available?lat=37.788989&lng=-122.404810&radius=20.0

Response:
[
    {
        "id": 10,
        "lat": 37.788548,
        "lng": -122.411548
    },
    {
        "id": 8,
        "lat": 37.783223,
        "lng": -122.398630
    }
]

2.) POST /api/v1/scooters/reserve?id=10

Once you reserve scooter 10, if you call the scooters available api (first example endpoint), it should not return scooter 10 as it is reserved.

