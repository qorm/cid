class CID {
    constructor() {
        this.keys = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    }
    generate() {
        let pad = Math.ceil(Math.random() * 4294967295)
        let now = new Date().getTime();
        return this.ten_to_u36(now,9) + this.ten_to_u36(pad,7)
    }
    getTime(id){
        return new Date(this.getTimestamp(id))
    }
    getTimestamp(id){
        return this.u36_to_ten(id.substr(0,9));
    }
    ten_to_u36(n,u) {
        let r = "";
        do {
            r = this.keys.charAt(n % 36) + r;
            n = parseInt(n / 36);
        } while (n != 0)

        for (let i = r.length; i < u; i++) {
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