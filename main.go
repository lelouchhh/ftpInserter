package main

import (
	"flag"
	"fmt"
	"ftpInserter/v1-beta.1/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/meilisearch/meilisearch-go"
	"time"
)

const (
	host     = "185.200.241.2"
	port     = 5432
	user     = "slave"
	password = ""
	dbname   = "dungeon"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "./", "path to the all XMLs that have to be inserted to database\nWARNING: PROGRAM WILL PARSE EVERY .XML FILE IN THE DIR")
	flag.Parse()
	fmt.Println("insert postgres")
	insertPostgres(path)
	fmt.Println("insert meilisearch")
	insertMeili(path)
}

func insertPostgres(path string) {
	insertEntity := "insert into purchase.t_entity (entity_id, purchase_number, status, description, date_start, date_end, date_start_auction, date_end_auction, original_url, name_etp, url_name_etp, fz, type, purchase_status, delivery_time, price, region, currency, amount_collateral, delivery_kladr_code, delivery_full_name, delivery_kladr, delivery_place,customer_phone, customer_email, customer_location, customer_org_full_name, customer_org_short_name, customer_reg_num, customer_contact_last_name, customer_contact_first_name, customer_contact_middle_name) values (:entity_id, :purchase_number, :status, :description, :date_start, :date_end, :date_start_auction , :date_end_auction, :original_url, :name_etp, :url_name_etp, :fz, :type, :purchase_status, :delivery_time, :price, :region, :currency, :amount_collateral, :delivery_kladr_code, :delivery_full_name, :delivery_kladr, :delivery_place,:customer_phone, :customer_email, :customer_location, :customer_org_full_name, :customer_org_short_name, :customer_reg_num, :customer_contact_last_name, :customer_contact_first_name, :customer_contact_middle_name)"
	insertDocs := "insert into purchase.t_document(document_id, document_name, document_url) values (:document_id, :document_name, :document_url)"
	insertLots := "insert into purchase.t_lot_item(sid, name, unit_price, sum, quantity, okpd_2_code, okei_title, code_mes, national_code_mes, name_mes) values (:sid, :name, :unit_price, :sum, :quantity, :okpd_2_code, :okei_title, :code_mes, :national_code_mes, :name_mes)"
	insertDocEntity := "insert into purchase.t_document_entity(document_id, entity_id) values (:document_id, :entity_id)"
	insertLotEntity := "insert into purchase.t_lot_entity(lot_id, entity_id) values (:lot_id, :entity_id)"
	insertDrugs := "insert into purchase.t_drug_lot_item(sid, med_form_name,dosage_grlsv_values ,dosage_code ,dosage_okei_name, dosage_user_name,drug_quantity ) values (:sid, :med_form_name,:dosage_grlsv_values ,:dosage_code ,:dosage_okei_name, :dosage_user_name,:drug_quantity)"
	insertDrugEntity := "insert into purchase.t_drug_entity(drug_id, entity_id ) values (:drug_id, :entity_id )"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	var filesJson []utils.FileJson
	var lotItems []utils.LotItems
	var drugLotItems []utils.DrugLotItems
	var documents []utils.Documents
	var docEntity []utils.DocEntity
	var lotEntity []utils.LotEntity
	var drugEntity []utils.DrugEntity
	var badFiles int
	var i int = -1
	bunch := 100
	start := time.Now()
	files, _ := utils.GetDir(path)
	fmt.Println(len(files))
	for _, file := range files {
		_, fileJson := utils.InsertInto(fmt.Sprintf("%s%s/%s", path, file.FullNameDir, file.File), file)
		for _, item := range fileJson.LotItems {
			lotItems = append(lotItems, item)
		}
		for _, item := range fileJson.Documents {
			documents = append(documents, item)
		}
		for _, item := range fileJson.DrugLotItems {
			drugLotItems = append(drugLotItems, item)
		}

		for _, item := range fileJson.Documents {
			var docEnt utils.DocEntity
			docEnt.EntityId = fileJson.Uid
			docEnt.DocumentId = item.Did
			docEntity = append(docEntity, docEnt)
		}
		for _, item := range fileJson.LotItems {
			var LotEnt utils.LotEntity
			LotEnt.EntityId = fileJson.Uid
			LotEnt.LotId = item.Sid
			lotEntity = append(lotEntity, LotEnt)
		}
		for _, item := range fileJson.DrugLotItems {
			var DrugEnt utils.DrugEntity
			DrugEnt.EntityId = fileJson.Uid
			DrugEnt.DrugId = item.Sid
			drugEntity = append(drugEntity, DrugEnt)
		}
	}

	lotItems = utils.UniqueLot(lotItems)
	documents = utils.UniqueDoc(documents)
	drugLotItems = utils.UniqueDrug(drugLotItems)

	for index, file := range files {
		_, fileJson := utils.InsertInto(fmt.Sprintf("%s%s/%s", path, file.FullNameDir, file.File), file)
		if fileJson.Uid != "" {
			filesJson = append(filesJson, fileJson)
			i++
		} else {

			badFiles++
		}
		if i%bunch == 0 && i != 0 {
			_, err := db.NamedExec(insertEntity, filesJson[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
		}
		if index == len(files)-1 {
			_, err := db.NamedExec(insertEntity, filesJson[i+1-(len(filesJson)%bunch):])
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Println("lot: ", len(lotItems))
	for i, _ := range lotItems {
		if i%bunch == 0 && i != 0 {
			_, err = db.NamedExec(insertLots, lotItems[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.NamedExec(insertLotEntity, lotEntity[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
		}
		if i == len(lotItems)-1 {
			_, err = db.NamedExec(insertLots, lotItems[i+1-(len(lotItems)%bunch):])
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.NamedExec(insertLotEntity, lotEntity[i+1-(len(lotItems)%bunch):])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("doc: ", len(documents), "docEntity", len(docEntity))
	for i, _ := range documents {
		if i%bunch == 0 && i != 0 {
			_, err = db.NamedExec(insertDocs, documents[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.NamedExec(insertDocEntity, docEntity[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
		}
		if i == len(documents)-1 {
			_, err = db.NamedExec(insertDocs, documents[i+1-(len(documents)%bunch):])
			if err != nil {
				fmt.Println(err)

			}
			_, err = db.NamedExec(insertDocEntity, docEntity[i+1-(len(documents)%bunch):])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("drug: ", len(drugLotItems))
	for i, _ := range drugLotItems {
		if i%bunch == 0 && i != 0 {
			_, err = db.NamedExec(insertDrugs, drugLotItems[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.NamedExec(insertDrugEntity, drugEntity[i-bunch:i])
			if err != nil {
				fmt.Println(err)
			}
		}
		if i == len(drugLotItems)-1 {
			_, err = db.NamedExec(insertDrugs, drugLotItems[i+1-(len(drugLotItems)%bunch):])
			if err != nil {
				fmt.Println(err)

			}
			_, err = db.NamedExec(insertDrugEntity, drugEntity[i+1-(len(drugLotItems)%bunch):])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println(i)
	fmt.Println(time.Since(start).Seconds())
}

func insertMeili(path string) {
	files, _ := utils.GetDir(path)
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "https://www.969975-cv27771.tmweb.ru:7700",
	})
	fmt.Println("files ", len(files))
	var filesJson []utils.FileJson
	var badFiles int
	var i int = -1
	bunch := 400
	start := time.Now()
	for index, file := range files {
		_, fileJson := utils.InsertInto(fmt.Sprintf("%s%s/%s", path, file.FullNameDir, file.File), file)
		if fileJson.Uid != "" {
			filesJson = append(filesJson, fileJson)
			i++
		} else {

			badFiles++
		}
		if i%bunch == 0 && i != 0 {
			client.Index("zakupkiGov").AddDocumentsInBatches(filesJson[i-bunch:i], bunch)
		}
		if index == len(files)-1 {
			client.Index("zakupkiGov").AddDocumentsInBatches(filesJson[i+1-(len(filesJson)%bunch):], bunch)
		}
	}
	fmt.Println(i)
	fmt.Println(time.Since(start).Seconds())
}
