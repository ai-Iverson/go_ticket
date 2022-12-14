// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go_ticket/internal/dao/internal"
)

// internalKnowledgeBaseDao is internal type for wrapping internal DAO implements.
type internalKnowledgeBaseDao = *internal.KnowledgeBaseDao

// knowledgeBaseDao is the data access object for table knowledge_base.
// You can define custom methods on it to extend its functionality as you wish.
type knowledgeBaseDao struct {
	internalKnowledgeBaseDao
}

var (
	// KnowledgeBase is globally public accessible object for table knowledge_base operations.
	KnowledgeBase = knowledgeBaseDao{
		internal.NewKnowledgeBaseDao(),
	}
)

// Fill with you ideas below.
