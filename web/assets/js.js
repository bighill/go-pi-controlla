let conn
const btns = document.getElementsByClassName( 'btn' )

const ws = () => {
    if (!window['WebSocket']) return wsError()

    conn            = new WebSocket('ws://' + document.location.host + '/ws')
    conn.onclose    = handleConnClose
    conn.onmessage  = handleConnMessage
}

const wsError = () => {
    document
        .getElementsByTagName('body')[0]
        .innerHTML = 'this browser does not support web sockets'
}

const handleConnClose   = ()        => console.log('ws connection closed')
const handleConnMessage = message   => console.log(message)

const btnEventListeners = btns => {
    for (btn of btns) btnEventListener(btn)
}

const btnEventListener = btn => {
    btn.addEventListener('mousedown', handleDown)
    btn.addEventListener('mouseup', handleUp)
}

const handleDown = ev => conn.send(ev.target.dataset.down)
const handleUp = ev => conn.send(ev.target.dataset.up)

const init = () => {
    ws()
    btnEventListeners(btns)

    window.oncontextmenu = function(event) {
     event.preventDefault();
     event.stopPropagation();
     return false;
};
}

init()
