package main

const getNodeData = `
{
  get(func: uid(0xfffd8d6ab2de70ec)) {
      data {
        id
        name
        data{
          number
          result
        }
        class
        html
        typenode
        inputs{
          input_1{
              connections{
                node
                input
            }
          }
          input_2{
            connections{
                node
                input
            }
          }
        }
        outputs{
          output_1{
            connections{
                  node
                  output
            }
          }
          output_2{
            connections{
              node
              input
            }
          }
        }
        pos_x
        pos_y
     }
   }
  }
`

// func getAllPrograms(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	dgClient := newClient()
// 	fmt.Print(dgClient)
// 	txn := dgClient.NewTxn()

// 	resp, err := txn.Query(context.Background(), getUser)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Write(resp.Json)
// }

// type Connections struct {
// 	Node string `json:"node"`
// 	Uid  string `json:"uid"`
// }

// func setAllPrograms(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	var p Connections
// 	newjson := json.NewDecoder(r.Body).Decode(&p)
// 	dgClient := newClient()
// 	txn := dgClient.NewTxn()

// 	fmt.Printf("body", r.Body)
// 	inputJson, err := json.Marshal(newjson)

// 	mu := &api.Mutation{
// 		SetJson:   inputJson,
// 		CommitNow: true,
// 	}
// 	log.Println("mutation object:", mu)

// 	resp, err := txn.Mutate(context.Background(), mu)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Write(resp.Json)
// }
