document.getElementById("dates").classList.add("active");

    var start_input = document.getElementById("start");
    var end_input = document.getElementById("end");

    var start_date = document.getElementById("start-date");
    var end_date = document.getElementById("end-date");

    var isStart = true;

    // change value of input dates for start value
    // prettier-ignore
    document.getElementById("old-time-start").addEventListener("click", () => {
        isStart = true
      start_input.value = "1970-01-01";
      change_dates(isStart)
    });
    // prettier-ignore
    document.getElementById("reset-time-start").addEventListener("click", () => {
        isStart = true
      start_input.value = "2017-01-01";
      change_dates(isStart)
    });
    // prettier-ignore
    document.getElementById("recent-time-start").addEventListener("click", () => {
        isStart = true
      start_input.value = "2022-01-01";
      change_dates(isStart)
    });

    // change value of input dates for end value
    var old_end = document.getElementById("old-time-end");
    old_end.addEventListener("click", () => {
      end_input.value = "1970-01-01";
      isStart = false;
      change_dates(isStart);
    });
    var reset_end = document.getElementById("reset-time-end");
    reset_end.addEventListener("click", () => {
      end_input.value = "2019-01-01";
      isStart = false;
      change_dates(isStart);
    });
    var recent_end = document.getElementById("recent-time-end");
    recent_end.addEventListener("click", () => {
      end_input.value = "2022-01-01";
      isStart = false;
      change_dates(isStart);
    });

    // change dates value written
    document.body.addEventListener("mouseover", () => {
        isStart = true
      change_dates(isStart)
      isStart = false
      change_dates(isStart)
   });

    function change_dates(isStart) {
      if (isStart) {
        var date = new Date(start_input.value);
      } else {
        var date = new Date(end_input.value);
      }
      // Calcule du jour
      day = date.getDate().toString();
      if (day.length == 1) {
        day = "0" + day;
      }
      // Calcule du Mois
      // prettier-ignore
      all_Month = ["January", "Febuary", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
      month = date.getMonth();
      for (let i = 0; i < all_Month.length; i++) {
        if (i == month) {
          month = all_Month[i];
        }
      }
      if (isStart) {
        start_date.innerHTML = day + " " + month + " " + date.getFullYear();
      } else {
        end_date.innerHTML = day + " " + month + " " + date.getFullYear();
      }
    }