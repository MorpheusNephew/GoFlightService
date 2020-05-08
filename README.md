# Flight Finder - Service

This project was inspired by the Coronavirus. At the time of this writing I am currently in Medellin, Colombia waiting this thing out. I have an Airbnb until the end of May and if things haven't gotten better and Medellin will be under quarantine for another 6 months to a year, direct quote from the president of Colombia, then I might as well come back to the states. I've signed up for the [Smart Traveler Enrollement Program (STEP)](STEP) so that I get all communications from the U.S. Embassy here in Colombia. There have been humanitarian flights out of Colombia, primarily from Bogota, back to Ft Lauderdale. The communications about these humanitarian flights come in the form of an email stating that Spirit has some flights going out. Sign-up for a seat online on the Spirit website. Due to being inundated with emails, basically I don't check them, I figured it would be cool to use tech to solve my problem of not being aware early enough to see when these flights are so I can book a ticket back home. This project is for the service portion of the project.

## How it will work

This project can be broken down into 3 main components:

- Status endpoint to ensure that it's actually running

- Cron job that checks Spirit's website to see if there are any available flights, this month or the next, from Medellin, Colombia to MIA/Ft Lauderdale

- If there is an available flight send that information to a lambda function to do the rest

[STEP]: https://step.state.gov/step/
