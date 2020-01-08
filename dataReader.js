function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

var infos = [];

for(var i = 0; i < 240; i++) {
  const c = document.querySelector("#\\38 0117716 > div.player-info > textarea");
  infos.push(c.value + "ThisIsADelimiter");
  await sleep(1000);
}

document.write(infos)
