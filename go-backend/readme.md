# this is backend created using golang + PostgreSQL (ent)

## to host this backend
1. Install go 
2. Get URA API access token from 
3. Setup ENV variable (.env file in `go-backend/.env`)
```
GOOGLE_CLIENT_ID=  // from google cloud API
GOOGLE_CLIENT_SECRET= // from google cloud API
############################## this key to be generated everyday
URA_API_TOKEN= // to be generated using URA_ACCESS_KEY
##############################
URA_ACCESS_KEY= // from URA API website 
GITHUB_KEY=
GITHUB_SECRET=
```
4. `go run main.go`
5. Example of API: `localhost:8080/search?location=Raffles&priceRange=0-3000&houseType=Non-landed%20Properties&size=30&numberOfRooms=1`


### Example
```javascript
import axios from 'axios'
import { useState, useEffect } from 'react';

const baseURL = "http://localhost:8080/search"

function App() {
	const [data, setData] = useState(null);
	
	useEffect(() => {
		// edit search param here
		axios.get(baseURL + '?location=Raffles&priceRange=0-3000&houseType=Non-landed%20Properties&size=30&numberOfRooms=1').then((res) => {setData(res.data);})
	}, [])

	if (data) {
		console.log(data)
		return <div>{JSON.stringify(data)}</div>
	} else {
		return <div>no data found</div>
	}
}

export default App;
```


## Not yet finished
- [X] User Account Creationg & Connect PostgreSQL with authentication token
- [X] User-PropertyListing Favorite Relation
