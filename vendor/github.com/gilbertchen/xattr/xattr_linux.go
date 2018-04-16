package xattr

import (
	"strings"
)

const (
	userPrefix = "user."
)

// Linux xattrs have a manditory prefix of "user.". This is prepended
// transparently for Get/Set/Remove and hidden in List

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, error) {
	name = userPrefix + name
	// find size.
	size, err := getxattr(path, name, nil, 0)
	//Typically if an xatter doesn't exist its not an error.
	//It is the same as it being empty.
	if err != nil && size < 0 {
		return nil, &XAttrError{"getxattr", path, name, err}
	}
	//Well it is, buts it's more polite to just say nothing.
	if size == 0 {
		return nil, nil
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, err := getxattr(path, name, &buf[0], size)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, name, err}
	}
	return buf[:read], nil
}

// Retrieves a list of names of extended attributes associated with the
// given path in the file system.
func Listxattr(path string) ([]string, error) {
	// find size.
	size, err := listxattr(path, nil, 0)
	//It's okay for the size to be 0, a file can have no xattr
	if err != nil && size < 0 {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
	//However if we don't have any xattr there no point buffing
	if size == 0 {
		return []string{}, nil
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, err := listxattr(path, &buf[0], size)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
	return stripUserPrefix(nullTermToStrings(buf[:read])), nil
}

// Associates name and data together as an attribute of path.
func Setxattr(path, name string, data []byte) error {
        length := len(data)
        if length == 0 {
            data = []byte(" ")
        }
	name = userPrefix + name
	if err := setxattr(path, name, &data[0], length); err != nil {
		return &XAttrError{"setxattr", path, name, err}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) error {
	name = userPrefix + name
	if err := removexattr(path, name); err != nil {
		return &XAttrError{"removexattr", path, name, err}
	}
	return nil
}

// Strip off "user." prefixes from attribute names.
func stripUserPrefix(s []string) []string {
	for i, a := range s {
		if strings.HasPrefix(a, userPrefix) {
			s[i] = a[5:]
		}
	}
	return s
}
