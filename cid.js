class CID {
    constructor() {
        this.keys = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    }
    generate() {
        let pad = Math.ceil(Math.random() * 4294967295)
        let now = new Date().getTime();
        return this.ten_to_u36(now) + this.ten_to_u36(pad)
    }
    getTime(id){
        let stamp = this.u36_to_ten(id.substr(0,8));
        return new Date(stamp)
    }
    ten_to_u36(n) {
        let r = "";
        do {
            r = this.keys.charAt(n % 36) + r;
            n = parseInt(n / 36);
        } while (n != 0)

        for (let i = r.length; i < 8; i++) {
            r = "0" + r
        }
        return r
    }
    u36_to_ten(s){
        let n = 0;
        let l = s.length;
        for (let i = 0;i<l;i++){
            let r = this.keys.indexOf(s.charAt(i))
            n = n + r * Math.pow(36,l-i-1)
        }
        return n;
    }
}