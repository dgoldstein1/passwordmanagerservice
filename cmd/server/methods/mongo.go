// mongo.go

package methods

import (
	"github.com/spf13/viper"
	"github.com/rs/zerolog"
	"github.com/globalsign/mgo"
	"errors"
	"time"
)


/**
 * Attempts to connect to mongodb using viper configuration variables
 **/
func ConnectToMongo(logger zerolog.Logger) (*mgo.Session, error) {
	logger.Debug().Msg("Connecting to mongodb at " + viper.GetString("mongodb_endpoint"))
	return mgo.DialWithTimeout(viper.GetString("mongodb_endpoint"), time.Duration(viper.GetInt("mongodb_timeout")) * time.Second)
}

/**
 * helper for copying mongo session for concurrency
 * @param {*mgo.Session}
 * @param {string collectionName}
 * @return {*mgo.Collection, *mgo.Session} for spaces
 **/
func CopySessionAndGetCollection(sess *mgo.Session, collectionName string) (*mgo.Collection, *mgo.Session, error) {
	if sess == nil {
		return nil, nil, errors.New("nil mongo session")
	}
	sessCopy := sess.Copy()
	spaces := sess.DB(viper.GetString("mongodb_name")).C(collectionName)
	return spaces, sessCopy, nil
}
