const API_URL=import.meta.env.VITE_API_URL



async function senddata(userreview,id){
const response=await fetch(API_URL + "/movies/" + id + "/tracking",{
    method:"PUT",
    headers:{
        "Content-Type":"application/json"
    },
    body:JSON.stringify(userreview),
})

if(!response.ok){
    throw new Error("could not save movie data")

}

const savedData=await response.json()
return savedData;
}

export {senddata}