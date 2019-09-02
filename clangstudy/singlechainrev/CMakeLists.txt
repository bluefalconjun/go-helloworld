#find all sources in current dir and set to var DIR_LIB_SRCS
aux_source_directory(. DIR_LIB_SRCS)

#generate lib 
#static lib
add_library (singlechainrev ${DIR_LIB_SRCS})
