package downtimes

// func get_all_downtimes_duplicate_checker(target_id int, target_scope string) ( error ){// -> List[Tuple[str, str, str]] | None:
//     URL str = "https://api.ddog-gov.com/api/v2/downtime"
//     HEADERS map[str] str] = headers.get(keys.api(), keys.app())
//     response = requests.get(url==URL, headers==HEADERS, stream==True)
//     if response.status_code != http.StatusOk{
//         fmt.Sprintf("Error: Unknown error:%v, %s",response.status_code, response.text)
//         print(response.text)
//         return None
// 	}
//     // result = response.json()
// 		// result_data, ok := make(map[string]string)
//
//
// 	//        {
// 	//            "monitor_id": item["attributes"]["monitor_identifier"]["monitor_id"],
// 	//            "scope": item["attributes"]["scope"],
// 	//        }
// 	//        for item in result["data"]
// 	//    ]
// 	// if err != nil{
// 	// 	fmt.Println(err)
// 	// }
// 	// return err
// 	// }
//
// // func add_downtime_one_scope(scope: int, monitor_id: Optional[int] = None) -> None:
// //     delta: int = 7200
// //     now = dt.now(timezone.utc) + timedelta(seconds=3)
// //     end = (now + timedelta(seconds=delta)).strftime("%Y-%m-%dT%H:%M:%S.%f")[:-3] + "Z"
// //     URL: str = "https://api.ddog-gov.com/api/v2/downtime"
// //     HEADERS: Dict[str, str] = headers.post(keys.api(), keys.app())
// //
// //     DATA = {
// //         "data": {
// //             "attributes": {
// //                 "message": f"Downtime for {monitor_id}",
// //                 "monitor_identifier": {"monitor_id": monitor_id},
// //                 "scope": scope,
// //                 "schedule": {
// //                     "start": None,
// //                     "end": end,
// //                 },
// //             },
// //             "type": "downtime",
// //         }
// //     }
// //     try:
// //         response = requests.post(url=URL, headers=HEADERS, json=DATA)
// //         if response.status_code != 200:
// //             print(f"Error: {response.status_code}")
// //             print(f"DOWNTIME NOT CREATED FOR {monitor_id}:")
// //             print(response.text)
// //             return
// //         result = response.json()
// //         pretty = json.dumps(result, indent=4)
// //         print(pretty)
// //     except Exception as e:
// //         print(f"Failed to run post request: {e}")
