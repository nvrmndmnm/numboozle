let startTime;
let currentNumber = 1;
let timerInterval;

function updateTimer() {
    if (startTime) {
        const elapsedTime = (Date.now() - startTime) / 1000;
        document.getElementById("timer").innerText = `Time: ${elapsedTime.toFixed(1)} seconds`;
        
        return elapsedTime;
    }
    
    return 0;
}

document.addEventListener("htmx:configRequest", function (evt) {
    let clickedNumber = parseInt(evt.detail.parameters.number);

    let target = evt.detail.target;
    target.dataset.time = updateTimer() * 1000;

    evt.detail.parameters.time = target.dataset.time;

    if (clickedNumber !== currentNumber) {
        evt.preventDefault();
    }
});

document.addEventListener("htmx:afterRequest", function (evt) {
    if (!startTime) {
        startTime = Date.now();
        timerInterval = setInterval(updateTimer, 100);
    }

    let target = evt.detail.target;
    let clickedNumber = parseInt(target.dataset.num);

    if (clickedNumber === currentNumber) {
        target.classList.add("cube-disappear");

        if (clickedNumber === 25) {
            clearInterval(timerInterval);
            let endTime = Date.now();
            let totalTime = (endTime - startTime) / 1000;

            document.getElementById("timer").innerText = `Time: ${totalTime.toFixed(2)} seconds`;

            startTime = null;
            currentNumber = 1;

            htmx.trigger("#game-area", "reload");
        } else {
            currentNumber++;
        }
    }
});
