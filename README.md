<h3>This is a sample code to initialize Context Logging</h3>
To Initialize Logger

`logger.InitializeLogger()`

Logger struct

`type contextLogger struct {
 	entry *logrus.Entry
 }`
 
 For every Request this `contextLogger` will be initialized
 
 `r.WithContext(logger.WithLogger(r.Context()))` in defaultmiddleware.go
 
 `func WithLogger(ctx context.Context) context.Context {
  	return context.WithValue(ctx, loggerKey{}, newContextLogger(ctx))
  }`
  
  <h4> To Use Logging anywhere<h4>
  
  `import logger "github.com/tandonraghav/go-logging/logging"`
  
  `logger.GetLogger(r.Context()).Info( "applyMid=",applyMid)`
  
  The Context is passed around and all Log lines will be 
  printed with Basic info like requestID, Username etc.
  
  <h4>Sample Log Lines</h4>
  
  `Jun  2 02:07:12.969 [/web/middlewares/defaultmiddleware.go 18]  [INFO] [201906020207M011559421432969213000] [NA] applyMid=true`
  
  `Jun  2 02:07:12.969 [/web/middlewares/defaultmiddleware.go 33]  [ERROR] [201906020207M011559421432969213000] [NA] Exception occured runtime error: integer divide by zero`