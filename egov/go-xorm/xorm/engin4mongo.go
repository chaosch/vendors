package xorm

//func (engine *Engine) Sync4Mongo(mgodb *mgo.Database, beans ...interface{}) error {
//	for _, bean := range beans {
//		v := rValue(bean)
//		tableName := engine.tbName(v)
//		table, err := engine.autoMapType4Mongo(mgodb, v)
//		s := engine.NewSession()
//		defer s.Close()
//
//		collections, err := mgodb.CollectionNames()
//
//		if err != nil {
//			fmt.Println(err)
//			return err
//		}
//
//		isExist := true
//		for _, name := range collections {
//			if name == tableName { //collection 存在
//				break
//			}
//		}
//		isExist = false
//
//		//isExist, err :=s.Table(bean).isTableExist(tableName)
//		if !isExist {
//			mgodb.C(tableName).Create(new(mgo.CollectionInfo))
//		}
//
//		//由于表结构基于数据库建立，不考虑下面的操作 by chaos
//		/*isEmpty, err := engine.IsEmptyTable(bean)
//		  if err != nil {
//		      return err
//		  }*/
//		var isEmpty = false
//		if isEmpty {
//		} else {
//			for _, index := range table.Indexes {
//				//session := engine.NewSession()
//				//defer session.Close()
//				//if err := session.Statement.setRefValue(v); err != nil {
//				//	return err
//				//}
//				if index.Type == core.UniqueType {
//					newindex := mgo.Index{}
//					newindex.Unique = true
//					newindex.Key = index.Cols
//					mgodb.C(tableName).EnsureIndex(newindex)
//				} else if index.Type == core.IndexType {
//					newindex := mgo.Index{}
//					newindex.Unique = false
//					newindex.Key = index.Cols
//					mgodb.C(tableName).EnsureIndex(newindex)
//				} else {
//					return errors.New("unknow index type")
//				}
//			}
//		}
//	}
//	return nil
//}
//
//func (engine *Engine) autoMapType4Mongo(mgodb *mgo.Database, v reflect.Value) (*core.Table, error) {
//	t := v.Type()
//	engine.mutex.Lock()
//	defer engine.mutex.Unlock()
//	table, ok := engine.Tables[t]
//	if !ok {
//		var err error
//		table, err = engine.mapType4Mongo(mgodb, v)
//
//		if err != nil {
//			return nil, err
//		}
//
//		engine.Tables[t] = table
//		if engine.Cacher != nil {
//			if v.CanAddr() {
//				engine.GobRegister(v.Addr().Interface())
//			} else {
//				engine.GobRegister(v.Interface())
//			}
//		}
//	}
//	return table, nil
//}
//
//func (engine *Engine) mapType4Mongo(mgodb *mgo.Database, v reflect.Value) (*core.Table, error) {
//	t := v.Type()
//	table := engine.newTable()
//	if tb, ok := v.Interface().(TableName); ok {
//		table.Name = tb.TableName()
//	} else {
//		if v.CanAddr() {
//			if tb, ok = v.Addr().Interface().(TableName); ok {
//				table.Name = tb.TableName()
//			}
//		}
//		if table.Name == "" {
//			table.Name = engine.TableMapper.Obj2Table(t.Name())
//		}
//	}
//
//	table.Type = t
//
//	var idFieldColName string
//	var hasCacheTag, hasNoCacheTag bool
//
//	for i := 0; i < t.NumField(); i++ {
//		tag := t.Field(i).Tag
//		//fmt.Println("tag:",tag)
//		ormTagStr := tag.Get(engine.TagIdentifier)
//		prefixTagStr := tag.Get(engine.TagStartWith)
//		commentTagStr := tag.Get(engine.TagComment)
//		//fmt.Println("TagIdentifier",engine.TagIdentifier)
//		//fmt.Println("ormTagStr:",ormTagStr)
//
//		//fmt.Println("prefixTagStr:",prefixTagStr)
//		//if (len(commentTagStr)>0){fmt.Println("commentTagStr:",commentTagStr)}
//		var col *core.Column
//		fieldValue := v.Field(i)
//		fieldType := fieldValue.Type()
//		if ormTagStr != "" {
//			col = &core.Column{FieldName: t.Field(i).Name, Nullable: true, IsPrimaryKey: false,
//				IsAutoIncrement: false, MapType: core.TWOSIDES, Indexes: make(map[string]int)}
//			tags := splitTag(ormTagStr)
//
//			if len(tags) > 0 {
//				if tags[0] == "-" {
//					continue
//				}
//
//				var ctx = tagContext{
//					table:      table,
//					col:        col,
//					fieldValue: fieldValue,
//					indexNames: make(map[string]int),
//					engine:     engine,
//				}
//
//				if strings.ToUpper(tags[0]) == "EXTENDS" {
//					if err := ExtendsTagHandler(&ctx); err != nil {
//						return nil, err
//					}
//					continue
//				}
//
//				for j, key := range tags {
//					if ctx.ignoreNext {
//						ctx.ignoreNext = false
//						continue
//					}
//					k := strings.ToUpper(key)
//					ctx.tagName = k
//					ctx.params = []string{}
//
//					pStart := strings.Index(k, "(")
//					if pStart == 0 {
//						return nil, errors.New("( could not be the first charactor")
//					}
//					if pStart > -1 {
//						if !strings.HasSuffix(k, ")") {
//							return nil, errors.New("cannot match ) charactor")
//						}
//
//						ctx.tagName = k[:pStart]
//						ctx.params = strings.Split(key[pStart+1:len(k)-1], ",")
//					}
//
//					if j > 0 {
//						ctx.preTag = strings.ToUpper(tags[j-1])
//					}
//					if j < len(tags)-1 {
//						ctx.nextTag = strings.ToUpper(tags[j+1])
//					} else {
//						ctx.nextTag = ""
//					}
//
//					if h, ok := engine.tagHandlers[ctx.tagName]; ok {
//						if err := h(&ctx); err != nil {
//							return nil, err
//						}
//					} else {
//						if strings.HasPrefix(key, "'") && strings.HasSuffix(key, "'") {
//							col.Name = key[1 : len(key)-1]
//						} else {
//							col.Name = key
//						}
//					}
//
//					if ctx.hasCacheTag {
//						hasCacheTag = true
//					}
//					if ctx.hasNoCacheTag {
//						hasNoCacheTag = true
//					}
//				}
//
//				if col.SQLType.Name == "" {
//					col.SQLType = core.Type2SQLType(fieldType)
//				}
//				engine.dialect.SqlType(col)
//				if col.Length == 0 {
//					col.Length = col.SQLType.DefaultLength
//				}
//				if col.Length2 == 0 {
//					col.Length2 = col.SQLType.DefaultLength2
//				}
//				if col.Name == "" {
//					col.Name = engine.ColumnMapper.Obj2Table(t.Field(i).Name)
//				}
//
//				if ctx.isUnique {
//					ctx.indexNames[col.Name] = core.UniqueType
//				} else if ctx.isIndex {
//					ctx.indexNames[col.Name] = core.IndexType
//				}
//
//				for indexName, indexType := range ctx.indexNames {
//					addIndex(indexName, table, col, indexType)
//				}
//
//			}
//		} else {
//			var sqlType core.SQLType
//			if fieldValue.CanAddr() {
//				if _, ok := fieldValue.Addr().Interface().(core.Conversion); ok {
//					sqlType = core.SQLType{Name: core.Text}
//				}
//			}
//			if _, ok := fieldValue.Interface().(core.Conversion); ok {
//				sqlType = core.SQLType{Name: core.Text}
//			} else {
//				sqlType = core.Type2SQLType(fieldType)
//			}
//			col = core.NewColumn(engine.ColumnMapper.Obj2Table(t.Field(i).Name),
//				t.Field(i).Name, sqlType, sqlType.DefaultLength,
//				sqlType.DefaultLength2, true)
//		}
//
//		if commentTagStr != "" {
//			col.Comment = commentTagStr
//		}
//
//		if col.IsAutoIncrement {
//			col.Nullable = false
//			if len(prefixTagStr) == 0 {
//				col.StartWith = 1
//			} else {
//				col.StartWith, _ = strconv.ParseInt(prefixTagStr, 10, 64)
//				if col.StartWith == 0 {
//					//指定初始值为0时设置为前缀
//					col.StartWith = core.Idt.Idprefix + 1
//				}
//			}
//		}
//
//		table.AddColumn(col)
//
//		if fieldType.Kind() == reflect.Int64 && (strings.ToUpper(col.FieldName) == "ID" || strings.HasSuffix(strings.ToUpper(col.FieldName), ".ID")) {
//			idFieldColName = col.Name
//		}
//	} // end for
//
//	if idFieldColName != "" && len(table.PrimaryKeys) == 0 {
//		col := table.GetColumn(idFieldColName)
//		col.IsPrimaryKey = true
//		col.IsAutoIncrement = true
//		col.Nullable = false
//		table.PrimaryKeys = append(table.PrimaryKeys, col.Name)
//		table.AutoIncrement = col.Name
//	}
//
//	if len(table.PrimaryKeys) > 0 { //建立唯一索引
//		newidex := mgo.Index{}
//		newidex.Key = table.PrimaryKeys
//		newidex.Unique = true
//		mgodb.C(table.Name).EnsureIndex(newidex)
//	}
//	if hasCacheTag {
//		if engine.Cacher != nil { // !nash! use engine's cacher if provided
//			engine.logger.Info("enable cache on table:", table.Name)
//			table.Cacher = engine.Cacher
//		} else {
//			engine.logger.Info("enable LRU cache on table:", table.Name)
//			table.Cacher = NewLRUCacher2(NewMemoryStore(), time.Hour, 10000) // !nashtsai! HACK use LRU cacher for now
//		}
//	}
//	if hasNoCacheTag {
//		engine.logger.Info("no cache on table:", table.Name)
//		table.Cacher = nil
//	}
//
//	return table, nil
//}
