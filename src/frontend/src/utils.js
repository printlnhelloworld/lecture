function formatDateYMD(now) {
  now = (now.toString().length === 13) ? now / 1000 : now;
  let time = new Date(parseInt(now) * 1000)
  let year = time.getFullYear();
  let month = time.getMonth() + 1;
  let date = time.getDate();
  return year + '年' + month + '月' + date + '日';
}
function formatDateHM(now) {
  now = (now.toString().length === 13) ? now / 1000 : now;
  let time = new Date(parseInt(now) * 1000)
  let hour = time.getHours();
  let minute = time.getMinutes();
  // let second = now.getSeconds();
  return hour + ':' + minute;
}
function formatDate(now) {
  let time;
  if (now instanceof Date) {
    time = now;
  } else {
    now = (now.toString().length === 13) ? now / 1000 : now;
    time = new Date(parseInt(now) * 1000);
  }
  let year = time.getFullYear();
  let month = time.getMonth() + 1;
  let date = time.getDate();
  let hour = time.getHours();
  let minute = time.getMinutes();
  // let second = now.getSeconds();
  return year + '/' + month + '/' + date + '/' + ' ' + hour + ':' + minute;
}
export { formatDateYMD, formatDateHM, formatDate }
// export { formatDateYMD, formatDateHM }
