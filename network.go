package main
import (
    "net/http"
    "github.com/GeoNet/web"
)


/**
*  get networks as json
*/
func networkH(w http.ResponseWriter, r *http.Request) {
    if len(r.URL.Query()) != 0 {
        web.BadRequest(w, r, "incorrect number of query params.")
        return
    }

    w.Header().Set("Content-Type", web.V1JSON)

    var d string

    err := db.QueryRow(
    `select row_to_json(fc) from (select array_to_json(array_agg(t)) as netwoek
        from (select networkid as "networkID", description from fits.network) as t) as fc`).Scan(&d)
    if err != nil {
        web.ServiceUnavailable(w, r, err)
        return
    }

    b := []byte(d)
    web.Ok(w, r, &b)
}
