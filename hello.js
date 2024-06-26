class Runoob {
    constructor(name, year) {
        this.name = name;
        this.year = year;
    }
    add(a,b) {
        return a + b;
    }
}

var runoob = new Runoob("菜鸟教程", 2018);
log(runoob.add(1, 6))