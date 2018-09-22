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

const btnEventListener = btn => btn.addEventListener('click', handleBtnClick)

const handleBtnClick = ev => conn.send(ev.target.dataset.do)

const init = () => {
    ws()
    btnEventListeners(btns)
}

init()
