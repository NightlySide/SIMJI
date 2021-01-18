function toastSuccess(title, message) {
    $("body").toast({
        title: title,
        class: "success",
        message: message,
        showProgress: "bottom",
        classProgress: "black",
        position: "bottom right",
        displayTime: 5000
    });
}