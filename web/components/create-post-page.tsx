"use client";

import type React from "react";

import { ChangeEvent, useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import {
  FileText,
  Video,
  Music,
  ImageIcon,
  Upload,
  Eye,
  Save,
  Send,
  X,
  Plus,
} from "lucide-react";

type PostType = "text" | "video" | "audio" | "image";

interface MediaFile {
  id: string;
  name: string;
  type: PostType;
  size: string;
  url?: string;
}

export function CreatePostPage() {
  const [selectedType, setSelectedType] = useState<PostType>("text");
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [mediaFiles, setMediaFiles] = useState<MediaFile[]>([]);
  const [tags, setTags] = useState<string[]>([]);
  const [newTag, setNewTag] = useState("");

  const postTypes = [
    { type: "text" as PostType, icon: FileText, label: "Text Post" },
    { type: "video" as PostType, icon: Video, label: "Video" },
    { type: "audio" as PostType, icon: Music, label: "Audio" },
    { type: "image" as PostType, icon: ImageIcon, label: "Image" },
  ];

  const handleFileUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files) {
      const newFiles: MediaFile[] = Array.from(files).map((file) => ({
        id: Math.random().toString(36).substr(2, 9),
        name: file.name,
        type: selectedType,
        size: `${(file.size / 1024 / 1024).toFixed(1)} MB`,
        url: URL.createObjectURL(file),
      }));
      setMediaFiles((prev) => [...prev, ...newFiles]);
    }
  };

  const removeMediaFile = (id: string) => {
    setMediaFiles((prev) => prev.filter((file) => file.id !== id));
  };

  const addTag = () => {
    if (newTag.trim() && !tags.includes(newTag.trim())) {
      setTags((prev) => [...prev, newTag.trim()]);
      setNewTag("");
    }
  };

  const removeTag = (tagToRemove: string) => {
    setTags((prev) => prev.filter((tag) => tag !== tagToRemove));
  };

  return (
    <div className="min-h-screen bg-background">
      <div className="container mx-auto px-4 py-8 max-w-6xl">
        <div className="mb-8">
          <h1 className="text-4xl font-bold text-foreground mb-2 text-balance">
            Create New Post
          </h1>
          <p className="text-muted-foreground text-lg">
            Share your content with your audience
          </p>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Main Content Area */}
          <div className="lg:col-span-2 space-y-6">
            <Card className="glass-card hover-glow">
              <CardHeader>
                <CardTitle className="text-xl font-bold">
                  Content Type
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                  {postTypes.map(({ type, icon: Icon, label }) => (
                    <button
                      key={type}
                      onClick={() => setSelectedType(type)}
                      className={`p-6 rounded-xl border-2 transition-all duration-300 hover:scale-105 group ${
                        selectedType === type
                          ? "border-primary bg-primary/20 neon-glow"
                          : "border-border hover:border-primary/50 glass-card"
                      }`}
                    >
                      <Icon
                        className={`w-8 h-8 mx-auto mb-3 transition-colors ${
                          selectedType === type
                            ? "text-primary"
                            : "text-muted-foreground group-hover:text-primary"
                        }`}
                      />
                      <span className="text-sm font-semibold block">
                        {label}
                      </span>
                    </button>
                  ))}
                </div>
              </CardContent>
            </Card>

            <Card className="glass-card hover-glow">
              <CardHeader>
                <CardTitle className="text-xl font-bold">
                  Post Details
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-6">
                <div>
                  <Label htmlFor="title" className="text-base font-semibold">
                    Title
                  </Label>
                  <Input
                    id="title"
                    placeholder="Give your post a compelling title..."
                    value={title}
                    onChange={(e: ChangeEvent<HTMLInputElement>) =>
                      setTitle(e.target.value)
                    }
                    className="mt-2 bg-input border-border focus:border-primary focus:neon-glow transition-all duration-300"
                  />
                </div>

                <div>
                  <Label htmlFor="content" className="text-base font-semibold">
                    Description
                  </Label>
                  <Textarea
                    id="content"
                    placeholder="Tell your audience about this post..."
                    value={content}
                    onChange={(e: ChangeEvent<HTMLTextAreaElement>) =>
                      setContent(e.target.value)
                    }
                    className="mt-2 min-h-32 bg-input border-border focus:border-primary focus:neon-glow transition-all duration-300"
                  />
                </div>

                <div>
                  <Label className="text-base font-semibold">Tags</Label>
                  <div className="flex flex-wrap gap-2 mt-3 mb-4">
                    {tags.map((tag) => (
                      <Badge
                        key={tag}
                        variant="secondary"
                        className="px-4 py-2 bg-accent/20 border border-accent/50 hover:bg-accent/30 transition-colors"
                      >
                        #{tag}
                        <button
                          onClick={() => removeTag(tag)}
                          className="ml-2 hover:text-destructive transition-colors"
                        >
                          <X className="w-3 h-3" />
                        </button>
                      </Badge>
                    ))}
                  </div>
                  <div className="flex gap-3">
                    <Input
                      placeholder="Add a tag..."
                      value={newTag}
                      onChange={(e: ChangeEvent<HTMLInputElement>) =>
                        setNewTag(e.target.value)
                      }
                      onKeyPress={(e: KeyboardEvent) =>
                        e.key === "Enter" && addTag()
                      }
                      className="flex-1 bg-input border-border focus:border-accent focus:neon-glow-purple transition-all duration-300"
                    />
                    <Button
                      onClick={addTag}
                      variant="outline"
                      size="sm"
                      className="hover:bg-accent/20 hover:border-accent transition-all duration-300 bg-transparent"
                    >
                      <Plus className="w-4 h-4" />
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>

            {selectedType !== "text" && (
              <Card className="glass-card hover-glow">
                <CardHeader>
                  <CardTitle className="text-xl font-bold">
                    Upload {selectedType}
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <div className="cyber-border rounded-xl p-8 text-center hover:neon-glow transition-all duration-300">
                    <Upload className="w-16 h-16 mx-auto mb-4 text-primary" />
                    <p className="text-xl font-semibold mb-2">
                      Drop your {selectedType} files here
                    </p>
                    <p className="text-muted-foreground mb-6">
                      or click to browse
                    </p>
                    <input
                      type="file"
                      multiple
                      accept={
                        selectedType === "video"
                          ? "video/*"
                          : selectedType === "audio"
                          ? "audio/*"
                          : selectedType === "image"
                          ? "image/*"
                          : "*/*"
                      }
                      onChange={handleFileUpload}
                      className="hidden"
                      id="file-upload"
                    />
                    <Button
                      asChild
                      variant="outline"
                      className="hover:bg-primary/20 hover:border-primary hover:neon-glow transition-all duration-300 bg-transparent"
                    >
                      <label htmlFor="file-upload" className="cursor-pointer">
                        Choose Files
                      </label>
                    </Button>
                  </div>

                  {mediaFiles.length > 0 && (
                    <div className="mt-6 space-y-3">
                      <h4 className="font-semibold text-lg">Uploaded Files</h4>
                      {mediaFiles.map((file) => (
                        <div
                          key={file.id}
                          className="flex items-center justify-between p-4 bg-muted/50 rounded-lg border border-border hover:border-primary/50 transition-all duration-300"
                        >
                          <div className="flex items-center gap-4">
                            {file.type === "video" && (
                              <Video className="w-6 h-6 text-accent" />
                            )}
                            {file.type === "audio" && (
                              <Music className="w-6 h-6 text-success" />
                            )}
                            {file.type === "image" && (
                              <ImageIcon className="w-6 h-6 text-primary" />
                            )}
                            <div>
                              <p className="font-medium">{file.name}</p>
                              <p className="text-sm text-muted-foreground">
                                {file.size}
                              </p>
                            </div>
                          </div>
                          <Button
                            onClick={() => removeMediaFile(file.id)}
                            variant="ghost"
                            size="sm"
                            className="hover:bg-destructive/20 hover:text-destructive transition-all duration-300"
                          >
                            <X className="w-4 h-4" />
                          </Button>
                        </div>
                      ))}
                    </div>
                  )}
                </CardContent>
              </Card>
            )}
          </div>

          <div className="space-y-6">
            {/* Preview */}
            <Card className="glass-card hover-glow">
              <CardHeader>
                <CardTitle className="text-xl font-bold flex items-center gap-3">
                  <Eye className="w-6 h-6 text-primary" />
                  Preview
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {title && (
                    <h3 className="font-bold text-xl text-balance">{title}</h3>
                  )}

                  {content && (
                    <p className="text-muted-foreground leading-relaxed">
                      {content}
                    </p>
                  )}

                  {tags.length > 0 && (
                    <div className="flex flex-wrap gap-2">
                      {tags.map((tag) => (
                        <Badge
                          key={tag}
                          variant="outline"
                          className="text-xs border-accent/50 text-accent"
                        >
                          #{tag}
                        </Badge>
                      ))}
                    </div>
                  )}

                  {mediaFiles.length > 0 && (
                    <div className="space-y-2">
                      <p className="text-sm font-medium text-primary">
                        {mediaFiles.length} file(s) attached
                      </p>
                      {mediaFiles.slice(0, 2).map((file) => (
                        <div
                          key={file.id}
                          className="text-sm p-3 bg-muted/30 rounded-lg border border-border"
                        >
                          {file.name}
                        </div>
                      ))}
                    </div>
                  )}

                  {!title && !content && mediaFiles.length === 0 && (
                    <p className="text-muted-foreground italic">
                      Your post preview will appear here...
                    </p>
                  )}
                </div>
              </CardContent>
            </Card>

            <Card className="glass-card hover-glow">
              <CardHeader>
                <CardTitle className="text-xl font-bold">
                  Publish Options
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <Button
                  className="w-full bg-primary hover:bg-primary/80 text-primary-foreground neon-glow hover:neon-glow transition-all duration-300"
                  size="lg"
                >
                  <Send className="w-5 h-5 mr-2" />
                  Publish Now
                </Button>

                <Button
                  variant="outline"
                  className="w-full bg-transparent border-accent text-accent hover:bg-accent/20 hover:neon-glow-purple transition-all duration-300"
                >
                  <Save className="w-5 h-5 mr-2" />
                  Save as Draft
                </Button>

                <Separator className="bg-border" />

                <div className="text-sm text-muted-foreground space-y-2 leading-relaxed">
                  <p>• Posts are visible to all your supporters</p>
                  <p>• You can edit or delete posts anytime</p>
                  <p>• Large files may take time to process</p>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
}
