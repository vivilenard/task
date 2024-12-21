const randomString = document.getElementById('randomString');

saveString();
getStoredString();

function saveString(){
    //  observing the paragraph for a change of content.
    //  if the content changes, save it in localStorage
    const observerCallback = (mutationsList, observer) => {
        for (const mutation of mutationsList) {
            if (mutation.type === 'childList') {
                localStorage.setItem("string", randomString.innerText);
            }}};
    const observer = new MutationObserver(observerCallback);
    const config = { childList: true };

    observer.observe(randomString, config);
}

async function getStoredString(){
    console.log('The page has been loaded or refreshed.');
    let stringContent = localStorage.getItem("string");
    // if nothing is saved ( like on first load ), just get a fresh one by calling the api
    if (!stringContent){
        let response = await fetch("http://127.0.0.1:3000/random")
        const arrayBuffer = await response.arrayBuffer();
        const decoder = new TextDecoder('utf-8');
        stringContent = decoder.decode(arrayBuffer);
        console.log('Decoded text:', stringContent);
    }
    randomString.innerText = stringContent;
}

