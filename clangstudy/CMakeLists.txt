#version require
cmake_minimum_required (VERSION 3.10)

#info
project (ctest)

#find all source and set to var DIR_SRCS
aux_source_directory(. DIR_SRCS)

#set exec
add_executable(ctest main.c)

include_directories("${PROJECT_SOURCE_DIR}/singlechainrev")
#add sub folder
add_subdirectory(singlechainrev)

#add linked lib
target_link_libraries(ctest singlechainrev)
