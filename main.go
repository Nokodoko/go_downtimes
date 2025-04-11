package main

import (
	"github.com/Nokodoko/go_downtimes/downtimes"
	"keys"
	"os"
)

// MONITOR_DICT := nil
// arg := os.Args[1]

func main(arg) {
	key, err := api("DD_API_KEY")
	switch {
	case err != nil:
		fmt.Println(err)
	default:
		fmt.Println(key)
	}
}

// func get_all_downtimes_duplicate_checker(target_id int, target_scope str) error{//r -> List[Tuple[str, str, str]] | None:
// 		URL: str = "https://api.ddog-gov.com/api/v2/downtime
// 		HEADERS: Dict[str, str] = headers.get(keys.api(), keys.app())
// 		response = requests.get(url=URL, headers=HEADERS, stream=True)
// 		if response.status_code != 200:
// 				print(f"Error: Unknown error: {response.status_code}\n {response.text}")
// 				print(response.text)
// 				return None
// 		result = response.json()
// 		result_data = [
// 				{
// 						"monitor_id": item["attributes"]["monitor_identifier"]["monitor_id"],
// 						"scope": item["attributes"]["scope"],
// 				}
// 				for item in result["data"]
// 		]
// 		return nil
// 	}
//
// func add_downtime_one_scope(scope: int, monitor_id: Optional[int] = None){
// 		delta: int = 7200
// 		now = dt.now(timezone.utc) + timedelta(seconds=3)
// 		end = (now + timedelta(seconds=delta)).strftime("%Y-%m-%dT%H:%M:%S.%f")[:-3] + "Z"
// 		URL: str = "https://api.ddog-gov.com/api/v2/downtime"
// 		HEADERS: Dict[str, str] = headers.post(keys.api(), keys.app())
//
// 		DATA = {
// 				"data": {
// 						"attributes": {
// 								"message": f"Downtime for {monitor_id}",
// 								"monitor_identifier": {"monitor_id": monitor_id},
// 								"scope": scope,
// 								"schedule": {
// 										"start": None,
// 										"end": end,
// 								},
// 						},
// 						"type": "downtime",
// 				}
// 		}
// 		try:
// 				response = requests.post(url=URL, headers=HEADERS, json=DATA)
// 				if response.status_code != 200:
// 						print(f"Error: {response.status_code}")
// 						print(f"DOWNTIME NOT CREATED FOR {monitor_id}:")
// 						print(response.text)
// 						return
// 				result = response.json()
// 				pretty = json.dumps(result, indent=4)
// 				print(pretty)
// 		except Exception as e:
// 				print(f"Failed to run post request: {e}")
// }
// }
