package main

const getNodeData = `
{
    get(func: has(nodesData)) {
        uid
        nodesData (orderasc: id){
            id
            name
            data {
                number
                result
                variable
                assign
                num1
                num2
                conditionResult
                option
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
