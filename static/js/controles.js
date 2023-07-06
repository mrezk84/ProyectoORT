$(document).ready(function() {
    getObrasByUser()
});




async function getObrasByUser() {
    const request = await fetch('control', {
        method: 'GET',
    })
    let response = await request.json()
    if (request.status == 200) {
        let text = ``;
        response.forEach(control => {
            text +=
                `
                 <tr>
                    <th>${control.id}</th>
                    <th>${control.descripcion}</th>
                    <th><button onclick="redirectEditarControl(${control.id})">Editar control</button></th>
                </tr>
                `
        });
        document.getElementById("controlesTBody").innerHTML = text;
    }
}

function redirectCrearControl(){
    window.location.href = `altaControl.html`;
}